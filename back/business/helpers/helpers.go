package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"url-shortener/business/users/models"
	"url-shortener/global"
	"url-shortener/utils/hash"
	"url-shortener/utils/jwt"
)

func CreatePassHashAndJWT(request *models.User) error {
	hashPass, err := hash.EncodePassword(request.Password)
	if err != nil {
		return err
	}
	request.Password = hashPass

	token, err := jwt.EncodeJWT([]byte(global.JWT_KEY), jwt.Claims{
		Pass: request.Password,
	})
	if err != nil {
		return err
	}
	request.Token = string(token)

	return nil
}

func VerifyPassword(request *models.User) (*models.User, error) {
	oldUser, err := models.GetUserWithEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(oldUser.Password),
		[]byte(request.Password)); err != nil {
		return nil, err
	}


	return oldUser, nil
}

func SetUserValuesFromOldUser(oldUser, user *models.User) {
	user.ID = oldUser.ID
	user.CreatedAt = oldUser.CreatedAt
	user.LastLogin = oldUser.LastLogin

	if user.Password == "" {
		user.Password = oldUser.Password
	}
	if user.Email == "" {
		user.Email = oldUser.Email
	}
	if user.Name == "" {
		user.Name = oldUser.Name
	}
}
