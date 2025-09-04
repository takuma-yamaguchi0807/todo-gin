package todo

import (
	"fmt"
)

type Status string

const (
	Wait  Status = "wait"
	Doing Status = "doing"
	Done  Status = "Done"
)

func NewStatus(v string) (Status, error) {
	switch Status(v) {
	case Wait,Doing,Done:
		return Status(v), nil
	default:
		return "", fmt.Errorf("invalid status: %s", v)
	}
}