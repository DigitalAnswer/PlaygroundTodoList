package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/DigitalAnswer/PlaygroundTodoList/helpers"
)

// NewAuthMiddleware func
func NewAuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			whiteList := []string{
				"/signin",
				"/signup",
			}

			if helpers.Contains(whiteList, r.URL.String()) {
				next.ServeHTTP(w, r)
				return
			}

			tokenString := r.Header.Get("token")
			if len(tokenString) <= 0 {
				helpers.FailureFromError(w, http.StatusUnauthorized, errors.New("Missing token"))
				return
			}

			claims, err := helpers.ParseJWT(tokenString)
			if err != nil {
				helpers.FailureFromError(w, http.StatusUnauthorized, err)
				return
			}
			ctx := context.WithValue(r.Context(), helpers.KeyPrincipalID, claims.UserID)
			r = r.WithContext(ctx)

			// Auth
			fmt.Printf("%s%s\n", r.Host, r.URL.String())
			next.ServeHTTP(w, r)
		})
	}
}
