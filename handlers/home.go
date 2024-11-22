package handlers

import (
    "blog/view/pages/home"
    "github.com/labstack/echo/v4"
    "net/http"
)

func registerHomeRoutes (e *echo.Echo) {
    e.GET("/", handleHome)
}

func handleHome(c echo.Context) error{
    return render(c, http.StatusOK, home.Home())
}
