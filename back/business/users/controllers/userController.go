package controllers

import (
	"github.com/kataras/iris/v12"
	"time"
	"url-shortener/business/helpers"
	"url-shortener/business/users/models"
	"url-shortener/global"
	"url-shortener/utils/json"
	"url-shortener/utils/jwt"
)

func GetUser(ctx iris.Context) {
	email := ctx.Params().GetString("email")

	user, err := models.GetUserWithEmail(email)
	if err != nil {
		json.WriteError(ctx, 400)
		return
	}

	json.WriteResponse(ctx, user)
}

func RegisterUser(ctx iris.Context) {
	request := &models.User{}
	if err := ctx.ReadJSON(request); err != nil {
		json.WriteError(ctx, 400)
		return
	}

	err := helpers.CreatePassHashAndJWT(request)
	if err != nil {
		json.WriteErrorWithMsg(ctx, 500, err.Error())
		return
	}

	/* check if email is duplicate */
	if _, err := models.GetUserWithEmail(request.Email); err == nil {
		json.WriteErrorWithMsg(ctx, 400, "User already exists!")
		return
	}

	response, err := request.Insert()
	if err != nil {
		json.WriteErrorWithMsg(ctx, 500, err.Error())
		return
	}

	json.WriteResponse(ctx, response)
}

func LoginUser(ctx iris.Context) {
	request := &models.User{}
	if err := ctx.ReadJSON(request); err != nil {
		json.WriteError(ctx, 400)
		return
	}

	oldUser, err := helpers.VerifyPassword(request)
	if err != nil {
		json.WriteError(ctx, 400)
		return
	}

	oldUser.LastLogin = time.Now()
	response, err := oldUser.Update()
	if err != nil {
		json.WriteErrorWithMsg(ctx, 500, err.Error())
		return
	}

	token, err := jwt.EncodeJWT([]byte(global.JWT_KEY), jwt.Claims{
		Pass: request.Password,
	})
	if err != nil {
		json.WriteError(ctx, 500)
		return
	}
	response.Token = string(token)

	json.WriteResponse(ctx, response)
}
