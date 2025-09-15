package user

import (
	"unicode"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/common"
)

// Password は平文パスワード（入力値）の値オブジェクト。
// 生成時に強度ポリシーを検証する。
// ハッシュ化は別責務（インフラ/サービス層）で行う。
type Password struct {
	value string
}

// NewPassword は長さと文字種ポリシーを検証して Password を生成する。
// ポリシー: 長さ8文字以上 かつ 次のうち2種類以上を含む
//  - 英大文字, 英小文字, 数字, 記号（その他）
func NewPassword(v string) (Password, error) {
	if len([]rune(v)) < 8 {
		return Password{}, common.InvalidErr("user.password", "password must be at least 8 characters", nil)
	}
	var hasUpper, hasLower, hasDigit, hasSymbol bool
	for _, r := range v {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		default:
			// 制御文字や空白も含むが、ここでは記号として許容
			hasSymbol = true
		}
	}
	var kinds int
	if hasUpper { kinds++ }
	if hasLower { kinds++ }
	if hasDigit { kinds++ }
	if hasSymbol { kinds++ }
	if kinds < 2 {
		return Password{}, common.InvalidErr("user.password", "password must include at least 2 of upper/lower/digit/symbol", nil)
	}
	return Password{value: v}, nil
}

// String は入力値を返す（ハッシュ前の利用に限定）。
func (p Password) String() string { return p.value }