package dto

type TodoCreateRequest struct {
	UserID      string  `json:"user_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	DueDate     *string `json:"due_date"`
}

type TodoCreateResponse struct {
	ID string `json:"id"`
}