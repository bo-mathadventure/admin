package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
)

func NewAdminAnnouncementHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminAnnouncement(ctx, db))
	app.Post("/", postAdminAnnouncement(ctx, db))
	app.Get("/:id", getAdminAnnouncementID(ctx, db))
	app.Put("/:id", putAdminAnnouncementID(ctx, db))
	app.Delete("/:id", deleteAdminAnnouncementID(ctx, db))
}

func getAdminAnnouncement(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func postAdminAnnouncement(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func getAdminAnnouncementID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func putAdminAnnouncementID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func deleteAdminAnnouncementID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}
