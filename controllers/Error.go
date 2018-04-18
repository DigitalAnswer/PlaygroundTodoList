package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

// ErrorMessageInterface interface
type ErrorMessageInterface interface {
	GetCode() int
	GetMessage() string
	Error() string
}

// ErrorMessage struct
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrorMessage) Error() string {
	return e.Message
}

func (e ErrorMessage) GetCode() int {
	return e.Code
}

func (e ErrorMessage) GetMessage() string {
	return e.Message
}

type ErrorResponse struct {
	Error ErrorMessage `json:"error"`
}

// FailureFromError write ErrorMessage from error
func FailureFromError(w http.ResponseWriter, status int, err error) {
	Failure(w, status, ErrorMessage{
		Code:    status,
		Message: err.Error(),
	})
}

// Failure response
func Failure(w http.ResponseWriter, status int, err ErrorMessage) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)

	body := ErrorResponse{
		Error: err,
	}

	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Error().Err(err).Msg("")
	}
}
