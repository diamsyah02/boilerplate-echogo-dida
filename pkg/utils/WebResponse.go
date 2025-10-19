package utils

type APIResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func WebResponse(statusCode int, message string, data interface{}) APIResponse {
	return APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}
