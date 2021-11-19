package routers

import (
	"github.com/kataras/iris/v12"
	"url-shortener/business/users/controllers"
)

const USER_ROUTE = "/user"

func setUserRoutes(app *iris.Application) {

	userRoute := app.Party(USER_ROUTE)
	/******************************  ROUTES  ******************************/
	//userRoute.Get("/register", controllers.GetUser)

	userRoute.Post("/register", controllers.RegisterUser)
	userRoute.Post("/login", controllers.LoginUser)
}
