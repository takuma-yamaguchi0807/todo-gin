package dto

type TodoUpdateRequest struct {
	id          string `json:"id"`
	userId      string `json:"user_id"`
	title       string `json:"title"`
	description string `json:"description"`
	status      string `json:"status"`
	dueDate     string `json:"due_date"`
}

type TodoUpdateResponse struct {
	
}
