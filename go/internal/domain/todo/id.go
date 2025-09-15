package todo

import (
	"github.com/google/uuid"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/common"
)

type Id struct {
    value uuid.UUID
}

func NewId(v string) (Id, error){
    id, err := uuid.Parse(v)
    if err != nil {
        return Id{}, common.InvalidErr("todo.id", "invalid uuid format")
    }
    return Id{value: id}, nil
}

// String は UUID を文字列で返します。
func (i Id) String() string { return i.value.String() }

// UUID は UUID 型の値を返します。
func (i Id) UUID() uuid.UUID { return i.value }
