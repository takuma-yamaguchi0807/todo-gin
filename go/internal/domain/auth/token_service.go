package auth

// TokenService はドメインの Claims を安全に「生成」および「検証」するための抽象です。
// 具体的なシリアライズ方式や署名方式（例: JWT HS256/RS256 など）には依存しません。
type TokenService interface {
	// Generate は与えられた Claims を署名済みのトークン文字列に変換します。
	Generate(Claims) (string, error)
	// Verify は与えられたトークン文字列を検証し、正しければドメインの Claims を返します。
	// 署名方式の不一致、期限切れ、改ざんが検出された場合はエラーを返します。
	Verify(token string) (Claims, error)
}


