package utils

type ErrorResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse(status string, message string) ErrorResponse {
	return ErrorResponse{
		Status: status,
		Message: message,
	}
}