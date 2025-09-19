## Go API (dev/prod 共通) - シンプル版

# 1) Build stage: バイナリをビルド
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# 小さく・移植性の高いバイナリを作る
RUN CGO_ENABLED=0 go build -o app ./cmd/api/main.go

# 2) Run stage: 軽量イメージで実行
FROM alpine:latest
RUN apk --no-cache add ca-certificates wget
WORKDIR /app
COPY --from=builder /app/app ./app
EXPOSE 8080
# 簡易ヘルスチェック（/healthz）
HEALTHCHECK CMD wget -qO /dev/null http://localhost:8080/healthz || exit 1
CMD ["./app"]
