package handlers

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
    registerHomeRoutes(e)
    registerAboutRoutes(e)
    registerBlogRoutes(e)
	registerPostRoutes(e)
}
