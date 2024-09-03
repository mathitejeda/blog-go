package handlers

import (
    "github.com/labstack/echo/v4"
    "blog/pages/home"
    "net/http"
)

func registerHomeRoutes (e *echo.Echo) {
    e.GET("/", handleHome)
}

func handleHome(c echo.Context) error{
    return render(c, http.StatusOK, home.Home("demonio"))
}
