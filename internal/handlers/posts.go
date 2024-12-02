package handlers

import (
	"blog/internal/view/pages/posts"
    "github.com/labstack/echo/v4"
    "net/http"
)

func registerPostRoutes (e *echo.Echo) {
	e.GET("/posts", handlePosts)
	e.GET("/posts/:id", handlePost)
}

func handlePosts (c echo.Context) error {
	return render(c, http.StatusOK, posts.Posts())
}

func handlePost (c echo.Context) error {
	return render(c, http.StatusOK, posts.Post())
}
