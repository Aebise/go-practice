package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("createapp"), nil
	})

	if err != nil {
		return err
	}

	return nil
}
