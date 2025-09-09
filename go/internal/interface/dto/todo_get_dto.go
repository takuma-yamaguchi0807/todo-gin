package dto

// TodoGetRequest はユーザーIDでTodoを検索するための入力DTO。
// 検証はドメイン層（user.Id）で行い、ここでは保持のみを行う。
type TodoGetRequest struct {
    UserID string `json:"user_id"`
}

type TodoGetResponse struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status"`
	DueDate     *string `json:"due_date,omitempty"`
}
