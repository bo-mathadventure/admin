package middleware

import (
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// JWTProtected return the middleware for JWT token checks
func JWTProtected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		TokenLookup: "header:X-API-Key",
		AuthScheme:  "",
		SigningKey:  []byte(config.GetConfig().WorkadventureSecretKey),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return handler.HandleErrorCode(c, fiber.StatusUnauthorized, "ERR_JWT_TOKEN")
		},
	})
}
