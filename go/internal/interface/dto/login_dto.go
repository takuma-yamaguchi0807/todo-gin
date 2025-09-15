package dto

// LoginRequest はログインの入力DTOです。
type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// LoginResponse はログインの出力DTOです。
// 最小構成ではアクセストークンのみ返却します。
type LoginResponse struct {
    AccessToken string `json:"access_token"`
}