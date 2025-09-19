package config

import (
	"database/sql"
	"fmt"

	// database/sql 用の Postgres ドライバ（pgx v5 の stdlib）
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DbConfig struct {
	DBDriver string
	Host     string
	Port     string
	User     string
	Pass     string
	Name     string
}

func Load() (*DbConfig, error) {
	// 環境変数から読み取り（一般的な命名 DB_XXXX を優先、未設定時は旧キー/デフォルト）
	host := GetenvOrDefault("DB_HOST", GetenvOrDefault("Host", "localhost"))
	port := GetenvOrDefault("DB_PORT", GetenvOrDefault("Port", "5432"))
	user := GetenvOrDefault("DB_USER", GetenvOrDefault("User", "postgres"))
	pass := GetenvOrDefault("DB_PASSWORD", GetenvOrDefault("Pass", "postgres"))
	name := GetenvOrDefault("DB_NAME", "todo")

	cfg := &DbConfig{
		DBDriver: GetenvOrDefault("DB_DRIVER", GetenvOrDefault("DBDriver", "pgx")),
		Host:     host,
		Port:     port,
		User:     user,
		Pass:     pass,
		Name:     name,
	}
	return cfg, nil
}

// OpenSQL は Postgres 固定で *sql.DB を返します。
// DB 名は最小構成として "todo" を利用します（必要に応じて拡張してください）。
func (c *DbConfig) OpenSQL() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Pass, c.Name,
	)
	// ドライバ名 "pgx" は stdlib パッケージのブランクインポートで登録済み
	return sql.Open("pgx", dsn)
}

func getenv(k string) string {
	// 後方互換のために残すが、GetenvOrDefault の利用を推奨
	if v, ok := Getenv(k); ok {
		return v
	}
	return ""
}
