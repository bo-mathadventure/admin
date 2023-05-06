package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
)

func NewAdminUserHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminUser(ctx, db))
	app.Post("/invite", postAdminUser(ctx, db))
	app.Get("/:id", getAdminUserID(ctx, db))
	app.Put("/:id", putAdminUserID(ctx, db))
	app.Delete("/:id", deleteAdminUserID(ctx, db))
}

func getAdminUser(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func postAdminUser(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func getAdminUserID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func putAdminUserID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func deleteAdminUserID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}
