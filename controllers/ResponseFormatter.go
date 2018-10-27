package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DigitalAnswer/PlaygroundTodoList/helpers"
)

// JSON response
func JSON(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		helpers.FailureFromError(w, http.StatusInternalServerError, err)
	}
}
