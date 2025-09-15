package auth

// TokenGenerator はドメインの Claims から署名済みトークンを生成します。
// 具体的なシリアライズ方式や署名方式には依存しません。
type TokenGenerator interface {
	Generate(Claims) (string, error)
}


