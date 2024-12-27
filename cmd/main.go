package main

import (
	"blog/internal/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/assets", "internal/view/static/assets")
	app.Static("public", "internal/view/public")
	app.File("/favicon.ico", "internal/view/public/favicon.ico")
	handlers.RegisterRoutes(app)
	app.Logger.Fatal(app.Start(":1323"))
}
