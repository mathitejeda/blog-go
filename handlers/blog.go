package handlers

import (
    "blog/view/pages/blog"
    "github.com/labstack/echo/v4"
    "net/http"
)

func registerBlogRoutes (e *echo.Echo) {
    e.GET("/blog", handleBlog)
}

func handleBlog (c echo.Context) error {
    return render(c, http.StatusOK, blog.Blog())
}

