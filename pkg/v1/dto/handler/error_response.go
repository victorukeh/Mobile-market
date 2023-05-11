package handler

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"message"`
}
