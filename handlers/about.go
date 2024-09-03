package handlers

import (
	"blog/pages/about"
	"github.com/labstack/echo/v4"
	"net/http"
)

func registerAboutRoutes(e *echo.Echo) {
	e.GET("/about", handleAbout)
}

func handleAbout (c echo.Context) error {
	return render(c, http.StatusOK, about.About())
}
