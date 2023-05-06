package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
)

func NewAdminReportHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminReport(ctx, db))
	app.Get("/:id", getAdminReportID(ctx, db))
	app.Put("/:id", putAdminReportID(ctx, db))
	app.Delete("/:id", deleteAdminReportID(ctx, db))
}

func getAdminReport(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func getAdminReportID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func putAdminReportID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}

func deleteAdminReportID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler.HandleSuccess(c)
	}
}
