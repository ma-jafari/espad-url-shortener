package controllers

import (
	"github.com/kataras/iris/v12"
	"url-shortener/business/url/models"
	"url-shortener/utils/hash"
	"url-shortener/utils/json"
)

func InsertUrl(ctx iris.Context) {
	request := &models.URL{}
	if err := ctx.ReadJSON(request); err != nil {
		json.WriteError(ctx, 400)
		return
	}

	request.ShortURL = hash.EncodeURL(request.OriginalURL)
	response, err := request.Insert()
	if err != nil {
		json.WriteErrorWithMsg(ctx, 500, err.Error())
		return
	}

	json.WriteResponse(ctx, response)
}

func GetUrl(ctx iris.Context) {
	hashURL := ctx.Params().GetString("string")

	//resultUrl, err := models.GetOriginalURLFromHash(hashURL)
	//if err != nil {
	//	json.WriteResponse(ctx, 400)
	//	return
	//}

	ctx.Redirect(hashURL, 301)
	return
}