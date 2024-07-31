package main

import (
    "net/http"
    "github.com/a-h/templ"
	"github.com/labstack/echo/v4"
    "blog/pages"
)

func main() {
    app := echo.New()
    app.GET("/", HomeHandler)
    app.Logger.Fatal(app.Start(":1323"))
}

func render(ctx echo.Context, statusCode int, t templ.Component) error {
    buffer := templ.GetBuffer()
    defer templ.ReleaseBuffer(buffer)

    if err := t.Render(ctx.Request().Context(), buffer); err != nil {
        return err
    }
    return ctx.HTML(statusCode, buffer.String())
}

func HomeHandler(c echo.Context) error {
    return render(c, http.StatusOK, pages.Home("Mathi"))
}
