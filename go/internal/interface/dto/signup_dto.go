package dto

// SignupRequest はユーザ登録の入力DTOです。
// email と password を受け取り、成功時はアクセストークンを返します。
type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignupResponse はユーザ登録の出力DTOです。
// 最小構成ではアクセストークンのみ返却します。
type SignupResponse struct {
	AccessToken string `json:"access_token"`
}
