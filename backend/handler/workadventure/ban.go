package workadventure

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/ban"
	"github.com/bo-mathadventure/admin/handler"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
	"time"
)

// NewBanHandler initialize routes for the given router
func NewBanHandler(ctx context.Context, app fiber.Router, db *ent.Client) {
	app.Get("/", getBan(ctx, db))
}

type banQuery struct {
	RoomURL   string `query:"roomUrl"`
	Token     string `query:"token"`
	IPAddress string `query:"ipAddress"`
}

func getBan(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		qData := new(banQuery)
		if err := c.QueryParser(qData); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if qData.RoomURL == "" || qData.Token == "" || qData.IPAddress == "" {
			return handler.HandleInsufficientData(c)
		}

		allBans, err := db.Ban.Query().Where(ban.Or(ban.CheckEQ(qData.Token), ban.CheckEQ(email.Normalize(qData.IPAddress)))).All(ctx)
		if err != nil {
			return c.SendStatus(404)
		}

		now := time.Now()
		var theBan *ent.Ban
		for _, b := range allBans {
			if b.ValidUntil.IsZero() {
				theBan = b
				break
			}
			if b.ValidUntil.After(now) {
				theBan = b
				break
			}
		}
		if theBan == nil {
			return c.SendStatus(404)
		}

		return c.JSON(map[string]interface{}{
			"is_banned": true,
			"message":   theBan.Message,
		})
	}
}
