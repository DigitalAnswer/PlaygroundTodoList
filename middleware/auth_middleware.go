package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"

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
			if len(tokenString) > 0 {
				parseJWT(tokenString)
			}
			// Auth
			fmt.Printf("%s%s\n", r.Host, r.URL.String())
			next.ServeHTTP(w, r)
		})
	}
}

func parseJWT(tokenString string) {

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("toto"), nil
	})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v\n", claims.Foo, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
}
