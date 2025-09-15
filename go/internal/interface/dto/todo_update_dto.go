package dto

type TodoUpdateRequest struct {
	ID          string  `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	DueDate     *string `json:"due_date"`
	UserID      string  `json:"user_id"`
}

type TodoUpdateResponse struct {
}
