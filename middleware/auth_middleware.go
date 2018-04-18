package middleware

import (
	"fmt"
	"net/http"
)

// NewAuthMiddleware func
func NewAuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Auth
			fmt.Printf("%s%s\n", r.Host, r.URL.String())
			next.ServeHTTP(w, r)
		})
	}
}
