package todo

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Title struct {
	value string
}

func NewTitle(v string) (Title, error) {
	v = strings.TrimSpace(v)
	length := utf8.RuneCountInString(v)
	if length == 0 || length > 100 {
		return Title{}, fmt.Errorf("title must be 1-100 chars")
	}
	return Title{value: v}, nil
}