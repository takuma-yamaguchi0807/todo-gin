package todo

import (
	"fmt"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/common"
)

type Status string

const (
	Wait  Status = "todo"
	Doing Status = "doing"
	Done  Status = "done"
)

func NewStatus(v string) (Status, error) {
	switch Status(v) {
	case Wait, Doing, Done:
		return Status(v), nil
	default:
		return "", common.InvalidErr("todo.status", fmt.Sprintf("invalid status: %s", v))
	}
}

// String はステータス文字列を返します。
func (s Status) String() string { return string(s) }
