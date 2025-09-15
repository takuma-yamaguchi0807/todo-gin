package dto

import "github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"

type TodoDetailRequest struct {
	ID string `json:"id"`
}

type TodoDetailResponse struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description *string `json:"description,omitempty"`
	Status      string `json:"status"`
	DueDate     *string `json:"due_date,omitempty"`
}

// NewTodoDetailResponseFromDomain はドメインモデルからレスポンスDTOへ変換する。
func NewTodoDetailResponseFromDomain(t todo.Todo) TodoDetailResponse {
	return TodoDetailResponse{
		ID:          t.ID().String(),
		UserID:      t.UserID().String(),
		Title:       t.Title().String(),
		Description: t.Description().Ptr(),
		Status:      t.Status().String(),
		DueDate:     t.DueDate().StringPtr(),
	}
}
