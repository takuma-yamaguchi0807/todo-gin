package usecase

import (
	"context"
	"strings"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/user"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/dto"
)

// TodoUpdateUsecase は更新系のユースケース。
// database/sql を前提とし、トランザクションは利用しない最小構成。
type TodoUpdateUsecase struct {
	repo todo.TodoRepository
}

func NewTodoUpdateUsecase(repo todo.TodoRepository) *TodoUpdateUsecase {
	return &TodoUpdateUsecase{repo: repo}
}

// Execute は1件の更新を行う（最小構成）。
// 仕様上は全項目更新（PUT）。部分更新が必要なら別ユースケースに切り出す。
func (uc *TodoUpdateUsecase) Execute(ctx context.Context, req dto.TodoUpdateRequest) error {
	// ここでは PUT 全更新として扱う。空文字は未設定として扱うため、VO へ入れる前に正規化。
	id, err := todo.NewId(req.ID)
	if err != nil {
		return err
	}

	// 既存は参照不要な最小構成として、リクエストから新しい値オブジェクトを組み立てる。
	// タイトルは必須だが、画面側で必須制御済み。サーバ側でもバリデート。
	var titleStr string
	if req.Title != nil {
		titleStr = *req.Title
	}
	title, err := todo.NewTitle(titleStr)
	if err != nil {
		return err
	}

	var descStr string
	if req.Description != nil {
		descStr = *req.Description
	}
	description, err := todo.NewDescription(descStr)
	if err != nil {
		return err
	}

	// ステータス既定値は現状 Wait。未指定時は既定に落とす。
	st := string(todo.Wait)
	if req.Status != nil && strings.TrimSpace(*req.Status) != "" {
		st = *req.Status
	}
	status, err := todo.NewStatus(st)
	if err != nil {
		return err
	}

	var dueStr string
	if req.DueDate != nil {
		dueStr = *req.DueDate
	}
	due, err := todo.NewDueDate(dueStr)
	if err != nil {
		return err
	}

	// 集約を作り直し（所有者はクレームから）
	uid, err := user.NewId(req.UserID)
	if err != nil {
		return err
	}
	t := todo.NewTodo(id, uid, title, description, status, due)

	return uc.repo.Update(ctx, &t)
}
