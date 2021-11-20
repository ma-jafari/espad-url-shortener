package helpers

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
	models2 "url-shortener/business/url/models"
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

func CheckAndValueSettingsPreInsertURL(request *models2.URL) error {
	if request.ShortURL == "" && request.UserID.Hex() != "" {
		request.ShortURL = hash.EncodeURL(request.UserID.Hex() + request.OriginalURL)
	} else if request.ShortURL == "" && !request.CreatedAt.IsZero() {
		request.ShortURL = hash.EncodeURL(request.CreatedAt.String() + request.OriginalURL)
	}
	if request.ExpireAt.IsZero() {
		request.ExpireAt = time.Now().AddDate(0, 0, models2.DEFAULT_EXPIRE_TIME)
	}
	_, err := models2.GetURLObjectFromShortURL(request.ShortURL)
	if err == nil {
		return errors.New("Duplicate short URL")
	}

	return nil
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
