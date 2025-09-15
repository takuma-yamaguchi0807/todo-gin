package dto

import (
	"github.com/takuma-yamaguchi0807/todo-gin/go/internal/domain/todo"
)


type TodoGetRequest struct {
    UserID string `json:"user_id"`
}

type TodoGetResponse struct {
    ID          string  `json:"id"`
    UserID      string  `json:"user_id"`
    Title       string  `json:"title"`
    Description *string `json:"description,omitempty"`
    Status      string `json:"status"`
    DueDate     *string `json:"due_date,omitempty"`
}

// NewTodoGetResponseFromDomain はドメインモデルからレスポンスDTOへ変換する。
func NewTodoGetResponseFromDomain(t todo.Todo) TodoGetResponse {
    idStr := t.ID().String()
    userStr := t.UserID().String()
    titleStr := t.Title().String()
    descPtr := t.Description().Ptr()
    statusStr := t.Status().String()
    duePtr := t.DueDate().StringPtr()
    return TodoGetResponse{
        ID:          idStr,
        UserID:      userStr,
        Title:       titleStr,
        Description: descPtr,
        Status:      statusStr,
        DueDate:     duePtr,
    }
}

// NewTodoGetResponseList はドメイン配列をレスポンスDTO配列へ変換する。
func NewTodoGetResponseList(items []*todo.Todo) []TodoGetResponse {
    res := make([]TodoGetResponse, 0, len(items))
    for _, t := range items {
        if t == nil { continue }
        res = append(res, NewTodoGetResponseFromDomain(*t))
    }
    return res
}
