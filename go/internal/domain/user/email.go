package user

import (
	"net/mail"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/common"
)

// Email はユーザのメールアドレスを表す値オブジェクト。
// 生成時に形式検証を行う。
type Email struct {
	value string
}

// NewEmail はメールアドレスの形式を検証し、Email を生成する。
// RFC に準拠した最低限の検証として net/mail を用い、
// 入力文字列がそのままアドレス部と一致することを確認する。
func NewEmail(v string) (Email, error) {
	addr, err := mail.ParseAddress(v)
	if err != nil || addr == nil || addr.Address != v {
		return Email{}, common.InvalidErr("user.email", "invalid email format", err)
	}
	return Email{value: v}, nil
}

// String はアドレス文字列を返す。
func (e Email) String() string { return e.value }