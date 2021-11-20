package controllers

import (
	"github.com/kataras/iris/v12"
	"time"
	"url-shortener/business/helpers"
	"url-shortener/business/url/models"
	"url-shortener/utils/json"
)

func InsertUrl(ctx iris.Context) {
	request := &models.URL{}
	if err := ctx.ReadJSON(request); err != nil {
		json.WriteError(ctx, 400)
		return
	}

	if err := helpers.CheckAndValueSettingsPreInsertURL(request); err != nil {
		json.WriteErrorWithMsg(ctx, 400, err.Error())
		return
	}

	response, err := request.Insert()
	if err != nil {
		json.WriteErrorWithMsg(ctx, 500, err.Error())
		return
	}

	json.WriteResponse(ctx, response)
}

func GetUrl(ctx iris.Context) {
	hashURL := ctx.Params().GetString("string")

	resultUrl, err := models.GetURLObjectFromShortURL(hashURL)
	if err != nil {
		json.WriteError(ctx, 404)
		return
	}

	if resultUrl.ExpireAt.Before(time.Now()) {
		json.WriteErrorWithMsg(ctx, 400, "Expired!")
		return
	}

	ctx.Redirect(resultUrl.OriginalURL, 301)
	return
}

func CheckUserURLSHistory(ctx iris.Context) {
	userID := ctx.Params().GetString("id")

	result, err := models.GetUserUrlHistory(userID)
	if err != nil {
		json.WriteError(ctx, 400)
		return
	}

	response := &models.URLHistories{}
	for _, url := range result {
		var isExpired bool
		if url.ExpireAt.Before(time.Now()) {
			isExpired = true
		} else {
			isExpired = false
		}

		urlHistory := &models.URLHistory{
			OriginalURL: url.OriginalURL,
			ShortURL:    url.ShortURL,
			IsExpired:   isExpired,
		}

		response.URLHistory = append(response.URLHistory, urlHistory)
	}

	json.WriteResponse(ctx, response)
}

func CheckShortURL(ctx iris.Context) {
	shortURL := ctx.Params().GetString("url")

	_, err := models.GetURLObjectFromShortURL(shortURL)
	if err == nil {
		json.WriteError(ctx, 400)
		return
	}

	json.WriteResponse(ctx, nil)
}
