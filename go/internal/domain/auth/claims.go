package auth

import "time"

// Claims はドメインが扱う認証上の最小情報です。
// 具体的なシリアライズ形式（JWTのsub/exp 等）には依存しません。
type Claims struct {
	UserID    string
	ExpiresAt time.Time
}

// NewClaims はTTLに基づいて失効時刻を設定したClaimsを返します。
// 既定のTTLは1時間です。
func NewClaims(userID string) Claims {
	return NewClaimsWithTTL(userID, time.Hour)
}

// NewClaimsWithTTL は任意のTTLでClaimsを生成します。
func NewClaimsWithTTL(userID string, ttl time.Duration) Claims {
	return Claims{
		UserID:    userID,
		ExpiresAt: time.Now().Add(ttl),
	}
}

// UserIDString はユーザIDを返します。
// 将来、ユーザIDの内部表現が変わっても呼び出し側を影響させないためのアクセサです。
func (c Claims) UserIDString() string { return c.UserID }