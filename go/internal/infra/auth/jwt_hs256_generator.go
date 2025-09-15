package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	domainauth "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/auth"
)

// HS256Generator はHS256で署名したJWTを生成する実装です。
type HS256Generator struct {
	secret []byte
	issuer string
}

func NewHS256Generator(secret string, issuer string) *HS256Generator {
	return &HS256Generator{secret: []byte(secret), issuer: issuer}
}

// jwtClaims はライブラリに渡すための内部表現。
type jwtClaims struct {
	sub string
	exp int64
	jwt.RegisteredClaims
}

// Generate はドメインの Claims を JWT に変換して署名します。
func (g *HS256Generator) Generate(c domainauth.Claims) (string, error) {
	if c.UserID == "" {
		return "", errors.New("empty user id")
	}
	exp := c.ExpiresAt
	if exp.IsZero() {
		exp = time.Now().Add(time.Hour)
	}
	claims := jwt.MapClaims{
		"sub": c.UserID,
		"exp": exp.Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// typ 固定（付与しなくても多くのクライアントは解釈可能だが、明示しておく）
	t.Header["typ"] = "JWT"
	// iss を RegisteredClaims で表現
	if g.issuer != "" {
		t.Claims.(jwt.MapClaims)["iss"] = g.issuer
	}
	return t.SignedString(g.secret)
}


