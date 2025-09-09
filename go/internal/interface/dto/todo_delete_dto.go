package dto

type TodoDeleteRequest struct {
	IDs []string `json:"ids"`
}

type TodoDeleteResponse struct {
}
