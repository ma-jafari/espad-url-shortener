package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

func Decode(key []byte, tokenString []byte) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(string(tokenString), claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}
