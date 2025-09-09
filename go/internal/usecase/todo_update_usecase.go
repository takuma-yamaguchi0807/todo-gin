package usecase

import (
    "context"

    "github.com/google/uuid"
    "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"
)

// TodoUpdateUsecase は更新系のユースケース。
// database/sql を前提とし、トランザクションは利用しない最小構成。
type TodoUpdateUsecase struct{
    repo todo.TodoRepository
}

func NewTodoUpdateUsecase(repo todo.TodoRepository) *TodoUpdateUsecase {
    return &TodoUpdateUsecase{repo: repo}
}

// Execute は1件の更新を行うダミー実装。
func (uc *TodoUpdateUsecase) Execute(ctx context.Context) error {
    id, _ := todo.NewId(uuid.NewString())
    return uc.repo.Update(ctx, id)
}
