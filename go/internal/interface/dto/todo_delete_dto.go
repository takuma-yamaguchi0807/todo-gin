package dto

type TodoDeleteRequest struct {
	IDs    []string `json:"ids"`
	UserID string   `json:"user_id"`
}

type TodoDeleteResponse struct {
}
