package usecase

import (
	"context"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/dto"
)

// TodoUpdateUsecase は更新系のユースケース。
// database/sql を前提とし、トランザクションは利用しない最小構成。
type TodoUpdateUsecase struct{
    repo todo.TodoRepository
}

func NewTodoUpdateUsecase(repo todo.TodoRepository) *TodoUpdateUsecase {
    return &TodoUpdateUsecase{repo: repo}
}

// Execute は1件の更新を行う（最小構成）。
// 仕様上は全項目更新（PUT）。部分更新が必要なら別ユースケースに切り出す。
func (uc *TodoUpdateUsecase) Execute(ctx context.Context, req dto.TodoUpdateRequest) error {
    // 所有者チェック・楽観ロック等は将来拡張。
    id, err := todo.NewId(req.ID)
    if err != nil { return err }
    return uc.repo.Update(ctx, id)
}
