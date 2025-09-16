# todo-gin

Go + Gin をバックエンド、TypeScript + React + Next をフロント、AWS をインフラ基盤として構築する学習用 TODO アプリの設計ドキュメントです。必要最小限の構成で、後から段階的に拡張できることを前提とします。

## 目的と読者

- 読者: 本リポジトリの開発者（バックエンド / フロント / インフラ）
- 目的: 最小構成の実装方針、API・画面仕様、インフラ構成、運用・監視を 1 ファイルで共有する
- 範囲: 認証・TODO の CRUD・最小の監視・最小の CI/CD

## 全体像（要約）

- バックエンド: Go 1.23 + Gin
- フロント: TypeScript + React + Next（SPA 静的配信）
- 認証: JWT（HS256）でログイン状態を管理（`sub` からユーザ ID を特定）
- インフラ: S3 + CloudFront（SPA 配信）、ALB + ECS Fargate（API）、RDS PostgreSQL（DB）、ECR（イメージ保管）
- CI/CD: GitHub Actions でビルド → ECR へ push → ECS を更新（詳細のジョブ例は割愛）
- オブザーバビリティ: Prometheus + Grafana（最小）

## アーキテクチャ図

```mermaid
graph TD
  UserBrowser -- HTTPS --> CloudFront
  CloudFront -- SPA --> S3
  UserBrowser -- HTTPS_API --> ALB
  ALB --> ECS_Fargate
  ECS_Fargate --> RDS_PostgreSQL
  ECS_Fargate --> Prometheus_Exporter
  Grafana --> Prometheus
  ECR --> ECS_Fargate
```

注記: 図のラベルは記号制約のため簡略化しています。実際は Route53 や ACM を併用します。

## ドメインモデル（最小）

- User
  - id: UUID
  - email: UNIQUE
  - password_hash
- Todo
  - id: UUID
  - user_id: FK -> User.id
  - title: string(<=140)
  - description?: string(<=2000)（任意）
  - status: enum("todo"|"doing"|"done") 既定: "todo"
  - due_date?: ISO8601（任意）

## テーブル定義（PostgreSQL）

```sql
-- users
CREATE TABLE users (
  id UUID PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password_hash TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- todos
CREATE TABLE todos (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  title TEXT NOT NULL CHECK (char_length(title) <= 140),
  description TEXT,
  status TEXT NOT NULL DEFAULT 'todo' CHECK (status IN ('todo','doing','done')),
  due_date TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_todos_user_id ON todos(user_id);
```

- `description` と `due_date` は任意（NULL 許可）
- アプリ側で UUID を生成して保存（DB ではデフォルト生成しない最小構成）

## API 仕様（JSON, 最小）

- ベース URL: 例 `https://api.example.com`（本番ドメインは別途設定）
- 認証: Bearer JWT（`Authorization: Bearer <token>`）
- CORS: 本番は SPA ドメインのみ許可。ローカルは `http://localhost:3000` 追加
- ユーザ特定: JWT の `sub` をユーザ ID として使用（`/me` は提供しない）

### Auth

- POST `/auth/signup`
  - 入力: `{ "email": string, "password": string }`
  - 出力: 201 返すだけ
- POST `/auth/login`
  - 入力: `{ "email": string, "password": string }`
  - 出力: `{ "access_token": string }`
- GET `/.well-known/jwks.json`（公開鍵の配布）

### Todos（認証必須・User スコープ）

- GET `/todos`
  - ログインユーザの Todo 一覧を返却（クエリパラメータなし）
- POST `/todos`
  - 入力: `{ "title": string, "description"?: string, "status"?: "todo"|"doing"|"done", "due_date"?: ISO8601 }`
  - 出力: 作成された Todo
- GET `/todos/{id}`
  - ログインユーザの所有物のみ
- PUT `/todos/{id}`（全項目更新; 最小構成）
  - 出力:200 返すだけ
- DELETE `/todos`
  - 入力: `{ "ids": string[] }`
  - 複数 ID を受け取り削除（単一削除も配列 1 要素で対応）。`DELETE` の JSON ボディを受け付けるポリシー
  - 出力:204 を返すだけ

ステータスコード: 2xx 成功、4xx バリデーション / 認可エラー、5xx サーバエラー

## 画面要件（最小）

- ログイン画面
  - email / password
- サインアップ画面
  - email / password（強度は最小）
- TODO 一覧
  - チェックボックス選択 → 一括削除、詳細リンク
- TODO 詳細
  - 表示 + 更新（タイトル等）
- TODO 登録
  - タイトル必須、他は任意

フロントは Next を静的化し S3 配信。API は別ドメイン（CORS 設定）

## 認証・セキュリティ（最小）

- JWT: HS256
  - 秘密鍵: Secrets Manager に保管
  - `sub`: ユーザ ID、`exp`: 有効期限（例 3600 秒）
  - 備考: 公開鍵/JWKS は RS256 採用時のみ必要（HS256 では不要）
- パスワード: bcrypt（コストはデフォルトから調整）
  - ポリシー: 最低 8 文字、かつ 次のうち 2 種類以上を含む（英大文字/英小文字/数字/記号）
- メールアドレス: RFC に準拠した形式検証（`net/mail` 相当の最低限チェック）
- CSRF: Authorization ヘッダで送るため不要（状態レス）
- CORS: 本番は SPA ドメインのみ許可

## オブザーバビリティ（最小）

- Prometheus: `/metrics` を公開（ECS タスクからスクレイプ可能なネットワークに配置）
- Grafana: ダッシュボードで可視化（Prometheus をデータソース）

## 必要な Docker コンテナ（ローカル想定）

- PostgreSQL: `postgres:16`
- Prometheus: `prom/prometheus`
- Grafana: `grafana/grafana`

PostgreSQL 起動例（参考）

```bash
docker run --name todo-pg -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=todo -p 5432:5432 -d postgres:16
```

Prometheus / Grafana は設定ファイルが必要なため、後日 `docker-compose.yml` を用意します（最小構成）。

## ローカル開発（最小）

前提

- Go 1.23+
- Docker（PostgreSQL 用）

環境変数（例）

```
# API
PORT=8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/todo?sslmode=disable

# JWT（ローカルは PEM ファイルパスでも可）
JWT_PRIVATE_KEY_PATH=./local/private.pem
JWT_PUBLIC_JWKS_PATH=./local/jwks.json
JWT_ISSUER=https://api.example.com
JWT_AUDIENCE=https://api.example.com
JWT_EXPIRES=3600
```

起動

```bash
cd go/cmd/api
go run .
# -> http://localhost:8080/healthz 等（実装に応じて）
```

## インフラ方針（最小）

- SPA: S3 配信 + CloudFront
- API: ALB → ECS Fargate（Go/Gin）→ RDS PostgreSQL
- 証明書/ドメイン: ACM + Route53
- 秘密情報: Secrets Manager（JWT 鍵、DB 接続）
- コンテナ: ECR に push し ECS タスク定義で参照

### CI/CD（最小）

- GitHub Actions でビルド → ECR へ push → ECS サービスを更新（ジョブ Yaml の例はこの設計書では省略）

## 実装メモ（シンプルさ最優先）

- 過剰な抽象化は避ける（まずはシンプルに）
- バリデーションは API 層で最小限（タイトル長など）
- リポジトリは最小のインターフェースで開始し、必要に応じて拡張
- トレースとメトリクスは最小から開始し、ボトルネック観測後に強化

## TODO（進捗管理）

- バックエンド（Go / Gin）

  - [x] Todos: 一覧取得（GET `/todos`）
  - [x] Todos: 詳細取得（GET `/todos/{id}`）
  - [x] Todos: 作成（POST `/todos`）
  - [x] Todos: 更新（PUT `/todos/{id}`）
  - [x] Todos: 一括削除（DELETE `/todos` with body `{ ids: string[] }`）
  - [x] login: ログイン（POST `/login`）
  - [x] singup: ユーザ登録（POST `/signup`）
  - [x] 認証: JWT 発行/検証（HS256）実装とミドルウェア適用
  - [x] エラーハンドリングの整理（エラー型と HTTP ステータスのマッピング）
  - [x] Prometheus メトリクス公開（`/metrics`）

- フロントエンド（Next.js）

  - [x] 初期セットアップ（TypeScript, ESLint/Prettier, ディレクトリ構成）
  - [ ] ホーム画面作成
  - [ ] 認証画面（ログイン/サインアップ）
  - [ ] TODO 一覧（取得・選択・一括削除）
  - [ ] TODO 詳細（表示・更新）
  - [ ] TODO 作成（タイトル必須、任意項目あり）
  - [ ] API クライアント（JWT 付与、エラー処理、型定義）

- インフラ（Terraform / AWS）

  - [ ] S3 + CloudFront（SPA 配信）
  - [ ] ECS Fargate + ALB（API）
  - [ ] RDS PostgreSQL（DB）
  - [ ] ECR（コンテナレジストリ）
  - [ ] Secrets Manager（JWT 秘密鍵、DB 接続）
  - [ ] Route53 + ACM（ドメイン/証明書）

- ローカル開発・運用
  - [ ] Docker コンテナ整備（PostgreSQL、Prometheus、Grafana、任意で OTel Collector）
  - [ ] docker-compose.yml の用意（最小構成）
  - [ ] GitHub Actions（ビルド →ECR push→ECS デプロイ）
  - [ ] README の手順更新（起動・環境変数・CI/CD 概要）
