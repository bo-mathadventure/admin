package workadventure

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/gofiber/fiber/v2"
)

func NewCapabilitiesHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/capabilities", getCaps(ctx, db))
}

func getCaps(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"api/woka/list":      "v1",
			"api/companion/list": "v1",
		})
	}
}
