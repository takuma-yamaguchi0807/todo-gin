package todo

import (
	"strings"
	"unicode/utf8"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/common"
)

type Description struct {
    value string
}

func NewDescription(v string) (Description, error) {
    v = strings.TrimSpace(v)
    length := utf8.RuneCountInString(v)
    if length > 300 {
        return Description{}, common.InvalidErr("todo.description", "description must be 0-300 chars")
    }
    return Description{value: v}, nil
}

// Ptr は空文字の場合に nil、そうでなければ *string を返します。
func (d Description) Ptr() *string {
    if d.value == "" {
        return nil
    }
    v := d.value
    return &v
}

// String は説明の文字列を返します（空は空文字）。
func (d Description) String() string { return d.value }
