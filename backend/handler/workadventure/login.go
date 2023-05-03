package workadventure

import (
	"context"
	"fmt"
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func NewLoginHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/login-url", getLoginURL(ctx, db))
}

type LoginURL struct {
	Token string `query:"token"`
}

func getLoginURL(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		qData := new(LoginURL)
		if err := c.QueryParser(qData); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if qData.Token == "" {
			return handler.HandleInsufficentData(c)
		}

		foundUser, err := db.User.Query().Where(user.TokenEQ(qData.Token)).First(ctx)
		if err != nil {
			return handler.HandleInvalidLogin(c)
		}

		url := ""
		if strings.HasPrefix(config.GetConfig().WorkadventureStartRoomURL, "/_/") {
			url = fmt.Sprintf("%s://%s", config.GetConfig().WorkadventureURLProtocol, config.GetConfig().WorkadventureStartRoomURL)
		} else if strings.HasPrefix(config.GetConfig().WorkadventureStartRoomURL, "/~/") {
			// mapurl without .tmj extension!
			url = fmt.Sprintf("%s%s", config.GetConfig().MapStorageURL, config.GetConfig().WorkadventureStartRoomURL)
		}

		return c.JSON(map[string]interface{}{
			"userUuid":    foundUser.UUID,
			"email":       foundUser.Email,
			"roomUrl":     config.GetConfig().WorkadventureStartRoomURL,
			"mapUrlStart": url,
			"messages":    []string{},
		})
	}
}
