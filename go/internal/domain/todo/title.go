package todo

import (
    "strings"
    "unicode/utf8"

    "github.com/takuma-yamaguchi0807/todo-gin/go/internal/app/apperror"
)

type Title struct {
    value string
}

func NewTitle(v string) (Title, error) {
    v = strings.TrimSpace(v)
    length := utf8.RuneCountInString(v)
    if length == 0 || length > 100 {
        return Title{}, apperror.InvalidErr("todo.title", "title must be 1-100 chars", nil)
    }
    return Title{value: v}, nil
}

// String はタイトルの文字列を返します。
func (t Title) String() string { return t.value }
