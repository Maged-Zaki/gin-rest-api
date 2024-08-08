package utils

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func FormatResponse(message string, data any) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}
