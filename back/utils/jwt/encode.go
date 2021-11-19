package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Pass string `json:"pass"`
	jwt.StandardClaims
}

func EncodeJWT(key []byte, claims Claims) ([]byte, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(key)

	return []byte(signedString), err
}
