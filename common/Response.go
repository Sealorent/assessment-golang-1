package common

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
}

func CreateResponse(success bool, message string, data any, err string) Response {
	return Response{
		Success: success,
		Message: message,
		Data:    data,
		Error:   err,
	}
}
