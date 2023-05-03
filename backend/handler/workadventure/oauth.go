package workadventure

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
)

func NewOAuthHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/logout", getLogout(ctx, db))
}

func getLogout(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}
