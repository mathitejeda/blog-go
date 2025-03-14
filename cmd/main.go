package main

import (
	"blog/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

func main() {
	app := echo.New()
	app.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	app.Static("/assets", "internal/view/static/assets")
	app.Static("public", "internal/view/public")
	app.File("/favicon.ico", "internal/view/public/favicon.ico")
	handlers.RegisterRoutes(app)
	app.Logger.Fatal(app.Start(":1323"))
}
