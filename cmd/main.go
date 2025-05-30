package main

import (
	"blog/internal/handlers"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	app.Static("/assets", "internal/view/static/assets")
	app.Static("public", "internal/view/public")
	handlers.RegisterRoutes(app)
	app.Logger.Fatal(app.Start(":8080"))
}
