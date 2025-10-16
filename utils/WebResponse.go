package utils

type WR struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

func WebResponse(statusCode int, message string, data interface{}, error interface{}) WR {
	return WR{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Error:      error,
	}
}
