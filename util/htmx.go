package util

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func RenderHtml(ctx *fiber.Ctx, c templ.Component) error {
	ctx.Set("Content-type", "text/html")
	return c.Render(ctx.Context(), ctx.Response().BodyWriter())
}
