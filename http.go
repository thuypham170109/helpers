package helpers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SendErrorResponse(code int, msg string, w http.ResponseWriter) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Message: msg,
		Data:    nil,
	})
}

func SendSuccessResponse(code int, msg string, data any, w http.ResponseWriter) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: msg,
		Data:    data,
	})
}
