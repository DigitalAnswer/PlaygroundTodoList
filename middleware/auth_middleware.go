package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/DigitalAnswer/PlaygroundTodoList/helpers"
	"github.com/dgrijalva/jwt-go"
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
			if len(tokenString) > 0 {
				if err := parseJWT(tokenString); err != nil {
					helpers.FailureFromError(w, http.StatusUnauthorized, err)
				}
			} else {
				helpers.FailureFromError(w, http.StatusUnauthorized, errors.New("Missing token"))
				return
			}

			// Auth
			fmt.Printf("%s%s\n", r.Host, r.URL.String())
			next.ServeHTTP(w, r)
		})
	}
}

func parseJWT(tokenString string) error {

	type MyCustomClaims struct {
		UserID       int64 `json:"user_id"`
		ExpDate      int64 `json:"exp"`
		CreationDate int64 `json:"iat"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("toto"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("userId: %v expiresAt: %v\n", claims.UserID, claims.ExpDate)
		return nil
	}

	fmt.Println(err)
	return err
}
