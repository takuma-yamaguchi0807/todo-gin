## フロントエンド (Next.js) 概要

本ディレクトリは Next.js (App Router, TypeScript) のフロントエンドです。最小構成から始め、必要に応じて機能や層を段階的に追加します。

### 必要環境

- Node.js 18+（推奨: LTS）
- npm（または pnpm / yarn でも可）

### セットアップ

1. 依存インストール

```bash
npm install
```

2. 開発サーバ起動

```bash
npm run dev
```

3. ブラウザでアクセス

```text
http://localhost:3000
```

### 主要スクリプト

- `npm run dev`: 開発サーバ起動
- `npm run build`: ビルド
- `npm run start`: 本番起動（ビルド後）
- `npm run lint`: ESLint 実行
- `npm run lint:fix`: 自動修正付き ESLint
- `npm run format`: Prettier で整形
- `npm run format:check`: 整形チェック

### ディレクトリ構成（最小）

```txt
src/
  app/
    layout.tsx       # 全ページ共通の枠と <meta>
    globals.css      # 全体スタイル（最小）
    page.tsx         # ルート (/)
    # 例: login/, signup/, todos/, todos/[id]/ などを順次追加
```

### 環境変数（例）

`next/.env.local`

```bash
NEXT_PUBLIC_API_BASE_URL=http://localhost:8080
```

- フロントから参照する変数は `NEXT_PUBLIC_` で始めてください。

### コーディング方針（最小・段階的拡張）

- まずはページ（`app/`）を作成し、画面遷移の全体像を確認
- 次に API クライアント（`fetch` 薄いラッパ）と型（DTO）を導入
- UI コンポーネントの共通化は必要になってから
- 状態管理ライブラリは不要な限り導入しない（複雑化時に検討）

### 補足

- App Router では `app/` のディレクトリ構成が URL になります（`[id]` は動的ルート）
- 生成物や依存は `.gitignore` で除外しています（`/.next`, `/node_modules`, `.env*` など）
