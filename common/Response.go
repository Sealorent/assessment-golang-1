package common

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
}

type ResponseSocialMedia struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    SocialMediaDTO `json:"data"`
	Error   string         `json:"error"`
}

type SocialMediaDTO struct {
	SocialMedia any `json:"social_medias"`
}

func CreateResponse(success bool, message string, data any, err string) Response {
	return Response{
		Success: success,
		Message: message,
		Data:    data,
		Error:   err,
	}
}
