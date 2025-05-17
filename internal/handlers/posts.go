package handlers

import (
	"blog/internal/model/post"
	"blog/internal/view/pages/posts"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
)

func registerPostRoutes(e *echo.Echo) {
	e.GET("/posts", handlePosts)
	e.GET("/posts/:id", handlePost)
}

func handlePosts(c echo.Context) error {
	postList := post.GetPosts()
	return render(c, http.StatusOK, posts.Posts(postList))
}

func handlePost(c echo.Context) error {
	id := c.Param("id")

	content, err := post.GetAsHTML(id)
	p := bluemonday.UGCPolicy()

	content = p.Sanitize(content)

	if err != nil {
		return c.String(http.StatusNotFound, "Post no encontrado")
	}

	return render(c, http.StatusOK, posts.Post(content))
}
