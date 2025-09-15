package auth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	domainauth "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/auth"
)

// HS256Generator はHS256で署名したJWTの「生成」と「検証」を行う実装です。
// 具体的に、Generate でドメインの Claims を JWT に変換して署名し、
// Verify で受け取ったJWT文字列の署名検証と有効期限・発行者を確認してドメイン Claims に復元します。
type HS256Generator struct {
	secret []byte
	issuer string
}

func NewHS256Generator(secret string, issuer string) *HS256Generator {
	return &HS256Generator{secret: []byte(secret), issuer: issuer}
}

// jwtClaims はライブラリに渡すための内部表現。（利用は MapClaims に統一）
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

// Verify はJWT文字列を検証し、正しければドメインの Claims に復元します。
// - 署名方式が HS256 であること
// - 署名が正しいこと
// - exp が有効であること（ライブラリの検証に依存）
// - issuer が設定されている場合は一致すること
func (g *HS256Generator) Verify(tokenStr string) (domainauth.Claims, error) {
	if tokenStr == "" {
		return domainauth.Claims{}, errors.New("empty token")
	}
	tok, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, errors.New("unexpected signing method")
		}
		return g.secret, nil
	})
	if err != nil {
		return domainauth.Claims{}, err
	}
	if !tok.Valid {
		return domainauth.Claims{}, errors.New("invalid token")
	}
	mc, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return domainauth.Claims{}, errors.New("invalid claims")
	}
	// issuer チェック（設定されている場合）
	if g.issuer != "" {
		iss, _ := mc["iss"].(string)
		if iss != g.issuer {
			return domainauth.Claims{}, errors.New("invalid issuer")
		}
	}
	// sub, exp をドメインの Claims にマッピング
	sub, _ := mc["sub"].(string)
	if sub == "" {
		return domainauth.Claims{}, errors.New("missing sub")
	}
	var expUnix int64
	switch v := mc["exp"].(type) {
	case float64:
		expUnix = int64(v)
	case int64:
		expUnix = v
	default:
	}
	if expUnix == 0 {
		// MapClaims.Valid() が検証しているが、ドメインにも反映できるよう best-effort で現在時刻+1h を設定
		expUnix = time.Now().Add(time.Hour).Unix()
	}
	return domainauth.Claims{
		UserID:    sub,
		ExpiresAt: time.Unix(expUnix, 0),
	}, nil
}


