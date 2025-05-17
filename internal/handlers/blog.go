package handlers

import (
	"blog/internal/model/post"
	"blog/internal/view/pages/blog"
	"github.com/labstack/echo/v4"
	"net/http"
)

func registerBlogRoutes(e *echo.Echo) {
	e.GET("/blog", handleBlog)
}

func handleBlog(c echo.Context) error {
	postList := post.GetPosts()
	return render(c, http.StatusOK, blog.Blog(postList))
}
