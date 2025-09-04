package todo

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Description struct {
	value string
}

func NewDescription(v string) (Description, error) {
	v = strings.TrimSpace(v)
	length := utf8.RuneCountInString(v)
	if length > 300 {
		return Description{}, fmt.Errorf("title must be 1-300 chars")
	}
	return Description{value: v}, nil
}