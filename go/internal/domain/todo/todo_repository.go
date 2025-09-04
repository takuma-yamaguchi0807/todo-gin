package todo

import (
	"context"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/user"
)

type TodoRepository interface {
	Save(ctx context.Context, t *Todo) error
	FindById(ctx context.Context, id Id) (*Todo, error)
	FindByUser(ctx context.Context, userId user.Id) ([]*Todo, error)
	Update(ctx context.Context, id Id) error
	DeleteByIds(ctx context.Context, ids [] Id) error
}