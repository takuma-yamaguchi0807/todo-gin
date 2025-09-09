package apperror

import "net/http"

// Kind はアプリケーションエラーの種別。
type Kind string

const (
    Invalid     Kind = "invalid"      // 入力値のバリデーションエラー（400）
    NotFound    Kind = "not_found"     // リソース未検出（404）
    Conflict    Kind = "conflict"      // 競合（409）
    Unauthorized Kind = "unauthorized" // 認証失敗（401）
    Forbidden   Kind = "forbidden"     // 権限不足（403）
    Internal    Kind = "internal"      // サーバ内部（500）
)

// Error はアプリケーションで用いる標準エラー。
type Error struct {
    Kind   Kind   // エラー種別
    Msg    string // 表示用メッセージ（クライアント向けに安全な内容）
    Field  string // 入力値に紐づくフィールド名（任意）
    cause  error  // 元エラー（任意）
}

func (e *Error) Error() string { return e.Msg }
func (e *Error) Unwrap() error { return e.cause }

// New は任意のエラーを生成する。
func New(k Kind, field, msg string, cause error) *Error {
    return &Error{Kind: k, Field: field, Msg: msg, cause: cause}
}

// Helpers
func InvalidErr(field, msg string, cause error) *Error { return New(Invalid, field, msg, cause) }
func NotFoundErr(resource, id string, cause error) *Error {
    return New(NotFound, resource, "resource not found", cause)
}
func ConflictErr(field, msg string, cause error) *Error { return New(Conflict, field, msg, cause) }

// StatusCode はHTTPステータスへマッピングする。
func StatusCode(err error) int {
    if e, ok := err.(*Error); ok {
        switch e.Kind {
        case Invalid:
            return http.StatusBadRequest
        case NotFound:
            return http.StatusNotFound
        case Conflict:
            return http.StatusConflict
        case Unauthorized:
            return http.StatusUnauthorized
        case Forbidden:
            return http.StatusForbidden
        default:
            return http.StatusInternalServerError
        }
    }
    return http.StatusInternalServerError
}
