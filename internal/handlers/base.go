package handlers

import (
    "github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(ctx echo.Context, statusCode int, t templ.Component) error {
    buffer := templ.GetBuffer()
    defer templ.ReleaseBuffer(buffer)

    if err := t.Render(ctx.Request().Context(), buffer); err != nil {
        return err
    }
    return ctx.HTML(statusCode, buffer.String())
}
