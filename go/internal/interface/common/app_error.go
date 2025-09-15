package common

import (
	"encoding/json"
	"net/http"
)

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

// StatusCode は HTTP ステータスコードのエイリアス。
// http.StatusXXX と互換性を保つため、int の型エイリアスとする。
type StatusCode = int

// Status は Kind から HTTP ステータスコードを得る。
func (k Kind) Status() StatusCode {
    switch k {
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

// Error はアプリケーションで用いる標準エラー。
type Error struct {
    Kind   Kind   // エラー種別
    Msg    string // 表示用メッセージ（クライアント向けに安全な内容）
    Field  string // 入力値に紐づくフィールド名（任意）
}

func (e *Error) Error() string { return e.Msg }

// New は任意のエラーを生成する。
func New(k Kind, field, msg string) *Error {
    return &Error{Kind: k, Field: field, Msg: msg}
}

// Helpers
func InvalidErr(field, msg string) *Error { return New(Invalid, field, msg) }
func NotFoundErr(resource, id string) *Error {
    return New(NotFound, resource, "resource not found")
}
func ConflictErr(field, msg string) *Error { return New(Conflict, field, msg) }

// StatusCode/StatusCodeByKind は Kind.Status へ集約したため削除

// JSON は Kind, Field, Msg を受け取り、
// *gin.Context.JSON にそのまま渡せる (status, payload) を返す。
func JSON(k Kind, field, msg string) (int, map[string]any) {
    payload := map[string]any{
        "error": string(k),
        "msg":   msg,
    }
    if field != "" {
        payload["field"] = field
    }
    return k.Status(), payload
}

// ToJSON はエラー情報と任意の追加データを JSON に変換する。
// 引数 data には、付加情報（任意）を渡せる。nil の場合は出力しない。
// 返却値は JSON のバイト列とエラー。
func (e *Error) ToJSON(data any) ([]byte, error) {
    payload := map[string]any{
        "error": string(e.Kind),
        "msg":   e.Msg,
    }
    if e.Field != "" {
        payload["field"] = e.Field
    }
    if data != nil {
        payload["data"] = data
    }
    return json.Marshal(payload)
}

// JSONFromError は任意の error から、標準化された (status, payload) を生成する。
// *Error の場合は Kind/Msg/Field を利用し、その他は internal エラーとして扱う。
func JSONFromError(err error) (int, map[string]any) {
    if err == nil {
        //エラーがないのに、呼ばれてるのは不正状態なので、internal エラーを返す
        return http.StatusInternalServerError, map[string]any{
            "error": string(Internal),
            "msg":   "internal server error",
        }
    }
    if e, ok := err.(*Error); ok {
        status := e.Kind.Status()
        payload := map[string]any{
            "error": string(e.Kind),
            "msg":   e.Msg,
        }
        if e.Field != "" {
            payload["field"] = e.Field
        }
        return status, payload
    }
    // 非アプリケーションエラーは internal に丸める
    return http.StatusInternalServerError, map[string]any{
        "error": string(Internal),
        "msg":   err.Error(),
    }
}
