package main

import (
	"github.com/labstack/echo/v4"
    "blog/handlers"
)

func main() {
    app := echo.New()
    handlers.RegisterRoutes(app)
    app.Logger.Fatal(app.Start(":1323"))
}



