package config

import "os"

// Getenv は環境変数の値を返します。存在しない場合は空文字と false を返します。
func Getenv(key string) (string, bool) {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return "", false
	}
	return v, true
}

// GetenvOrDefault は環境変数が未設定または空の場合、既定値を返します。
func GetenvOrDefault(key, def string) string {
	if v, ok := Getenv(key); ok {
		return v
	}
	return def
}

// MustGetenv は環境変数が未設定または空のときにパニックします。
func MustGetenv(key string) string {
	v, ok := Getenv(key)
	if !ok {
		panic("missing required env: " + key)
	}
	return v
}