package main

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"url-shortener/db"
	"url-shortener/routers"
)

func main() {
	app := iris.New()
	db.OpenConnection()
	routers.SetRoutes(app)

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.UseRouter(crs)

	app.Listen(":6010")
}
