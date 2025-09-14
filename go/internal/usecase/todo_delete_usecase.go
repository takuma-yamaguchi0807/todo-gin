package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/dto"
)

// TodoDeleteUsecase は削除系のユースケース。
// database/sql を前提とし、トランザクションは利用しない最小構成。
type TodoDeleteUsecase struct{
    repo todo.TodoRepository
}

func NewTodoDeleteUsecase(repo todo.TodoRepository) *TodoDeleteUsecase {
    return &TodoDeleteUsecase{repo: repo}
}

func (uc *TodoDeleteUsecase) Execute(ctx context.Context, req dto.TodoDeleteRequest) error {
    id, _ := todo.NewId(uuid.NewString())
    return uc.repo.DeleteByIds(ctx, []todo.Id{id})
}
