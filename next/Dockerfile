## Next.js (dev/prod 共通) - シンプル版

# 1) Build stage: 依存関係を入れてビルド
FROM node:20-alpine AS builder
WORKDIR /app
COPY package.json package-lock.json* ./
RUN npm ci
COPY . .
# Docker 最適化（必要なファイルを standalone に出力）
ENV NEXT_TELEMETRY_DISABLED=1
RUN npm run build

# 2) Run stage: 実行に必要な最小ファイルで起動
FROM node:20-alpine AS runner
WORKDIR /app
ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
EXPOSE 3000
HEALTHCHECK CMD wget -qO /dev/null http://localhost:3000/ || exit 1
# standalone 出力には server.js が含まれる
CMD ["node", "server.js"]
