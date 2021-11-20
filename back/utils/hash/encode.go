package hash

import (
	"crypto/md5"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

func EncodeURL(url string) string {
	hashedUrl := md5.Sum([]byte(url))
	byteUrl := hashedUrl[:]

	/* Return encoded url */
	return base64.StdEncoding.EncodeToString(byteUrl)[:8]
}

func EncodePassword(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPass), nil
}
