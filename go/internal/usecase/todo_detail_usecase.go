package usecase

import (
	"context"

	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/common"
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/interface/dto"
)

type TodoDetailUsecase struct {
	repo todo.TodoRepository
}

func NewTodoDetailUsecase(repo todo.TodoRepository) *TodoDetailUsecase {
	return &TodoDetailUsecase{repo: repo}
}

func (uc *TodoDetailUsecase) Execute(ctx context.Context, req dto.TodoDetailRequest) (dto.TodoDetailResponse, error) {
	id, err := todo.NewId(req.ID)
	if err != nil {
		return dto.TodoDetailResponse{}, err
	}
	t, err := uc.repo.FindById(ctx, id)
	if err != nil {
		return dto.TodoDetailResponse{}, common.NotFoundErr("todo", req.ID)
	}
	return dto.NewTodoDetailResponseFromDomain(*t), nil
}