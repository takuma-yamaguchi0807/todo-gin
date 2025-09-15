package config

import (
    "database/sql"
    "fmt"
    "os"

    // database/sql 用の Postgres ドライバ（pgx v5 の stdlib）
    _ "github.com/jackc/pgx/v5/stdlib"
)

type DbConfig struct {
	DBDriver string
	Host     string
	Port     string
	User     string
	Pass     string
}

func Load() (*DbConfig, error) {
    cfg := &DbConfig{
        DBDriver: getenv("DBDriver"),
        Host: getenv("Host"),
        Port: getenv("Port"),
        User: getenv("User"),
        Pass: getenv("Pass"),
    }
    return cfg,nil
}

// OpenSQL は Postgres 固定で *sql.DB を返します。
// DB 名は最小構成として "todo" を利用します（必要に応じて拡張してください）。
func (c *DbConfig) OpenSQL() (*sql.DB, error) {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        c.Host, c.Port, c.User, c.Pass, "todo",
    )
    // ドライバ名 "pgx" は stdlib パッケージのブランクインポートで登録済み
    return sql.Open("pgx", dsn)
}

func getenv(k string) string {
	v, ok := os.LookupEnv(k)
	if !ok || v == "" {
		panic(fmt.Sprintf("failed to read env: %s", k))
	}
	return v
}
