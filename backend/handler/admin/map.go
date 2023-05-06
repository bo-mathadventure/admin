package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
)

func NewAdminMapHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminMap(ctx, db))
	app.Post("/", postAdminMap(ctx, db))
	app.Get("/:id", getAdminMapID(ctx, db))
	app.Put("/:id", putAdminMapID(ctx, db))
	app.Delete("/:id", deleteAdminMapID(ctx, db))
}

func getAdminMap(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func postAdminMap(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func getAdminMapID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func putAdminMapID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func deleteAdminMapID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}
