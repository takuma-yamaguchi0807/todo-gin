package usecase

import (
    "context"

    "github.com/google/uuid"
    "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"
)

// TodoDeleteUsecase は削除系のユースケース。
// database/sql を前提とし、トランザクションは利用しない最小構成。
type TodoDeleteUsecase struct{
    repo todo.TodoRepository
}

func NewTodoDeleteUsecase(repo todo.TodoRepository) *TodoDeleteUsecase {
    return &TodoDeleteUsecase{repo: repo}
}

// Execute はID1件の削除を行うダミー実装。
func (uc *TodoDeleteUsecase) Execute(ctx context.Context) error {
    id, _ := todo.NewId(uuid.NewString())
    return uc.repo.DeleteByIds(ctx, []todo.Id{id})
}
