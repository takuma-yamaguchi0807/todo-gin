package usecase

import (
	"context"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/user"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/dto"
)

// TodoGetUsecase は参照系のユースケース。
// repository の結果（ドメイン）を DTO に詰め替えて返す。
type TodoGetUsecase struct{
    repo todo.TodoRepository
}

func NewTodoGetUsecase(repo todo.TodoRepository) *TodoGetUsecase {
    return &TodoGetUsecase{repo: repo}
}

// Execute はユーザーに紐づく一覧取得を行い、DTO を返却する。
// Execute は引数のリクエストに含まれる UserID に紐づく Todo 一覧を返します。
func (uc *TodoGetUsecase) Execute(ctx context.Context, req dto.TodoGetRequest) ([]dto.TodoGetResponse, error) {
    // UserID は JWT クレームから注入され、JSON には含めない
    uid, err := user.NewId(req.UserID)
    if err != nil {
        return []dto.TodoGetResponse{}, err
    }
    items, err := uc.repo.FindByUser(ctx, uid)
    if err != nil {
        return []dto.TodoGetResponse{}, err
    }

    return dto.NewTodoGetResponseList(items), nil
}
