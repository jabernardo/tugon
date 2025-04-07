package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SuccessResponse struct {
	Succes bool `json:"success"`
	Data   any  `json:"data"`
}

type FailureResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseOptions struct {
	StatusCode int
}

func NewSuccessResponse(data any) *SuccessResponse {
	return &SuccessResponse{Succes: true, Data: data}
}

func NewFailureResponse(code int, message string) *FailureResponse {
	return &FailureResponse{Success: false, Code: code, Message: message}
}

func (s *SuccessResponse) Write(w http.ResponseWriter, opts *ResponseOptions) {
	var statusCode int = 200

	if opts != nil {
		statusCode = opts.StatusCode
	}

	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(s)

	if err != nil {
		fmt.Fprintf(w, "{}")
	}

	w.Header().Set("Content-Type", "application/json")
}
