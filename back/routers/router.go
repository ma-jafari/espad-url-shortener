package routers

import (
	"github.com/kataras/iris/v12"
	"url-shortener/business/url/controllers"
)

func SetRoutes(app *iris.Application) {
	app.Party("/").Get("/{string}", controllers.GetUrl)
	setUrlRoutes(app)
	setUserRoutes(app)
}
