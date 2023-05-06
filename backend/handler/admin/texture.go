package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
)

func NewAdminTextureHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminTexture(ctx, db))
	app.Post("/", postAdminTexture(ctx, db))
	app.Get("/:id", getAdminTextureID(ctx, db))
	app.Put("/:id", putAdminTextureID(ctx, db))
	app.Delete("/:id", deleteAdminTextureID(ctx, db))
}

func getAdminTexture(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func postAdminTexture(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func getAdminTextureID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func putAdminTextureID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func deleteAdminTextureID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}
