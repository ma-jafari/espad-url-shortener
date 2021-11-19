package routers

import (
	"github.com/kataras/iris/v12"
	"url-shortener/business/url/controllers"
)

const URL_ROUTE = "/url"

func setUrlRoutes(app *iris.Application) {

	urlRoute := app.Party(URL_ROUTE)
	/******************************  ROUTES  ******************************/
	urlRoute.Post("", controllers.InsertUrl)
}
