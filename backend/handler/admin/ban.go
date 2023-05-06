package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
)

func NewAdminBanHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminBan(ctx, db))
	app.Post("/", postAdminBan(ctx, db))
	app.Get("/:id", getAdminBanID(ctx, db))
	app.Put("/:id", putAdminBanID(ctx, db))
	app.Delete("/:id", deleteAdminBanID(ctx, db))
}

func getAdminBan(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func postAdminBan(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func getAdminBanID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func putAdminBanID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func deleteAdminBanID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}
