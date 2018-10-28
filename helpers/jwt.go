package helpers

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// MyCustomClaims struct
type MyCustomClaims struct {
	UserID       int64 `json:"user_id"`
	ExpDate      int64 `json:"exp"`
	CreationDate int64 `json:"iat"`
	jwt.StandardClaims
}

func ParseJWT(tokenString string) (*MyCustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("toto"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		if claims.ExpDate > time.Now().Unix() {
			return claims, nil
		}
		return nil, errors.New("Token expired")
	}

	fmt.Println(err)
	return nil, err
}
