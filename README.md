# todo-gin

Go + Gin で実装する TODO アプリの学習用プロジェクトです。DDD 構成、JWT 認証（RS256）、SPA 配信（S3 + CloudFront）、API（ALB → ECS Fargate → RDS/PostgreSQL）、IaC（Terraform）を前提にします。

本 README は v1（初稿）です。実装/構成が固まり次第、随時更新します。

## スタック / 方針

- バックエンド: Go 1.23 + Gin（このリポジトリ）
- データアクセス: 生 SQL（`database/sql` ベース）。将来 `sqlc` や GORM への置き換え可能
- 認証/認可: 自前ユーザ DB + Bearer JWT（RS256）
  - サーバ自身が署名（RSA 秘密鍵）し、公開鍵は JWKS で配布
- フロント: Next.js（静的エクスポート中心）/ S3 から配信、API は別ドメイン
- インフラ（本番のみ）:
  - S3 + CloudFront（SPA 配信）
  - ALB → ECS Fargate（API）→ RDS (PostgreSQL)（シングル AZ / 最小構成）
  - Route53 + ACM、WAF、Secrets Manager（JWT 鍵や DB 接続情報）
  - ECR（API のコンテナイメージ保管）
- IaC: Terraform（ローカル state、ロック無し）
- ドメイン: `app.aws-traning.com`（SPA）/ `api.aws-traning.com`（API）
- CORS: 本番は `https://app.aws-traning.com` のみ許可。ローカルは `http://localhost:3000` を追加
- 削除方針: 物理削除から開始（必要に応じて将来論理削除に拡張）

補足: マイグレーションは `golang-migrate` を想定（一般的な選択肢）。代替として `goose`/`dbmate`/`atlas` 等もあり。

## ドメインモデル（DDD）

- User
  - id (UUID)
  - email (unique)
  - name
  - password_hash（自前認証時のみ）
  - created_at / updated_at
- Todo
  - id (UUID)
  - user_id (FK -> User.id)
  - title (<= 140)
  - description (<= 2000) 任意
  - status: enum(`todo`|`doing`|`done`), 既定: `todo`
  - due_date (nullable)
  - created_at / updated_at

## API 仕様（共通・JSON）

- ベース URL: `https://api.aws-traning.com`
- CORS: Origin `https://app.aws-traning.com` のみ許可（ローカルは `http://localhost:3000`）
- 認証: Bearer JWT（RS256, `Authorization: Bearer <token>`）

### Auth

- POST `/auth/login`
  - 入力: `{ "email": string, "password": string }`
  - 出力: `{ "access_token": string, "token_type": "Bearer", "expires_in": number, "user": { id, email, name } }`
  - パスワードは Argon2id（推奨）で検証（Bcrypt でも可）
- GET `/me`（認証必須）
  - 出力: `{ id, email, name }`
- GET `/.well-known/jwks.json`（公開鍵配布; 自前発行の検証用）

### Todos（認証必須・User スコープ）

- GET `/todos?status=&limit=&offset=`
  - ログインユーザの Todo のみ返却（`user_id = JWT.sub`）
- POST `/todos`
  - 入力: `{ "title": string(<=140), "description"?: string(<=2000), "status"?: "todo"|"doing"|"done", "due_date"?: ISO8601 }`
  - 出力: 作成された Todo
- GET `/todos/{id}`
  - ログインユーザの所有物のみ
- PUT `/todos/{id}`（全項目更新; PATCH でも可）
- DELETE `/todos/{id}`（最小は物理削除）

## アーキテクチャ

- SPA: Next.js を `npm run build && next export` で静的化 → `S3` へデプロイ → `CloudFront` 配信
- API: `ALB` → `ECS Fargate (Go/Gin)` → `RDS(PostgreSQL)`
- 証明書/ドメイン: `ACM`（リージョン適合に注意）+ `Route53`
- セキュリティ: `WAF`（CloudFront/ALB いずれか適用）、`Secrets Manager` に秘密情報保管
- コンテナ: `ECR` にイメージ push、ECS タスク定義で参照

## ローカル開発

前提

- Go 1.23+
- Node.js 18+（フロント別リポジトリ想定）
- Docker（PostgreSQL 用）

PostgreSQL（Docker 例）

```bash
docker run --name todo-pg -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=todo -p 5432:5432 -d postgres:16
```

環境変数（例）

```
# API
PORT=8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/todo?sslmode=disable

# JWT（ローカルは PEM 文字列 or ファイルパス指定のどちらかを選択）
JWT_PRIVATE_KEY_PATH=./local/private.pem
JWT_PUBLIC_JWKS_PATH=./local/jwks.json  # ローカル配布用（本番はHTTPで配布）
JWT_ISSUER=https://api.aws-traning.com
JWT_AUDIENCE=https://api.aws-traning.com
JWT_EXPIRES=3600
```

マイグレーション（予定: golang-migrate）

```
migrations/
  0001_init.up.sql
  0001_init.down.sql

# 例）適用
migrate -database "$DATABASE_URL" -path migrations up
```

API の起動（現在は最小のエンドポイントのみ）

```bash
cd go
go run ./main.go
# -> http://localhost:8080/healthz, /me
```

## JWT/鍵管理

- 署名: RS256（RSA）
- 本番: Secrets Manager に RSA 秘密鍵（PEM）を保管し、ECS タスク起動時に環境変数/ファイルとして渡す
- JWKS: `GET /.well-known/jwks.json` を API で配布（公開鍵）。フロントからの検証が不要でも、他サービス連携や将来用に用意
- 鍵ローテーション: `kid` 対応の JWKS を返却し、複数鍵の併存を許容

## Terraform（方針）

- 単一環境（prod のみ）・ローカル state（S3 Backend/ロック無し）
- 最小構成でコスト抑制（RDS: `db.t4g.micro`/Single-AZ, Fargate: 0.25vCPU/0.5GB）
- 構成例

```
infra/
  prod/
    main.tf
    variables.tf
    outputs.tf
    s3_cf/           # SPA 配信用 S3/CloudFront/ACM/Route53
    api_alb_ecs/     # ALB/ECS(Fargate)/ECR
    rds/             # PostgreSQL
    waf/
    secrets/
```

## セキュリティ/運用メモ

- CORS: 本番は `https://app.aws-traning.com` のみ許可。ローカルのみ `http://localhost:3000` を追加
- Cookie は使わず Authorization ヘッダで送信（XSS 考慮）。CSRF は不要（状態レス）
- パスワードハッシュ: Argon2id（メモリ/反復ハードニング）を推奨。Bcrypt を使う場合は cost 調整
- RDS への接続は最低限 SG/サブネット分離、公開しない。ALB→ECS のみ外向き
- WAF は最低限のレート制限/共通ルールセットを適用

## 進行状況

- [x] 技術方針の合意（Go+Gin、生 SQL、RS256、自前認証、単一 prod 環境）
- [x] README v1（この文書）
- [ ] DDD スケルトン配置（cmd/internal 整備）
- [ ] マイグレーション雛形（users, todos）
- [ ] 認証（/auth/login, /me, JWKS）
- [ ] Todos CRUD 実装
- [ ] Terraform（最小構成）

---

質問/未確定点

- パスワードハッシュは Argon2id で進めます（問題なければこのまま固めます）
- `sqlc` を使い「生 SQL + 型安全」にする案もあります（導入希望あれば対応）
- Terraform のリージョンは `ap-northeast-1` 前提でよいですか？
