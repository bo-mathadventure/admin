package workadventure

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/gofiber/fiber/v2"
)

// NewCapabilitiesHandler initialize routes for the given router
func NewCapabilitiesHandler(ctx context.Context, app fiber.Router, db *ent.Client) {
	app.Get("/", getCaps(ctx, db))
}

func getCaps(_ context.Context, _ *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"api/woka/list":      "v1",
			"api/companion/list": "v1",
		})
	}
}
