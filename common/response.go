package common

import "net/http"

type AppResponse struct{
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	Metadata string `json:"metadata"`
}

func NewFullResponse(statusCode int,message,metadata string) * AppResponse{
	return &AppResponse{
		StatusCode: statusCode,
		Message: message,
		Metadata: metadata,
	}
}
// 200
func StatusOkResponse(message,metadata string) *AppResponse{
	return &AppResponse{
		StatusCode: http.StatusOK,
		Message: message,
		Metadata: metadata,
	}
}
// 201
func StatusCreatedResponse(message,metadata string) *AppResponse{
	return &AppResponse{
		StatusCode: http.StatusCreated,
		Message: message,
		Metadata: metadata,
	}
}
