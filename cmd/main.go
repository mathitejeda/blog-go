package main

import (
	"blog/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/assets", "view/static/assets")
	app.Static("public", "view/public")
	app.File("/favicon.ico", "view/public/favicon.ico")
	handlers.RegisterRoutes(app)
	app.Logger.Fatal(app.Start(":1323"))
}
