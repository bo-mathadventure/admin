package middleware

import (
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
)

func AdminAPIProtected() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Get(fiber.HeaderAuthorization) != config.GetConfig().WorkadventureAdminAPISecret {
			return handler.HandleErrorCode(c, fiber.StatusUnauthorized, "ERR_API_TOKEN")
		}
		return c.Next()
	}
}
