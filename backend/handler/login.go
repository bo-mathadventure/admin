package handler

import (
	"context"
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/utils"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"time"
)

type LoginRequest struct {
	EMail             string `json:"email"`
	ClearTextPassword string `json:"password"`
}

func Login(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(LoginRequest)
		if err := c.BodyParser(req); err != nil {
			return HandleBodyParseError(c, err)
		}

		foundUser, err := db.User.Query().Where(user.Email(email.Normalize(req.EMail))).First(ctx)
		if err != nil || foundUser == nil || !utils.CheckPasswordHash(req.ClearTextPassword, foundUser.Password) {
			return HandleInvalidLogin(c)
		}

		// Create the Claims
		claims := jwt.MapClaims{
			"id":    foundUser.ID,
			"email": foundUser.Email,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(config.GetConfig().JWTSecret))
		if err != nil {
			return HandleInternalError(c, err)
		}

		log.WithFields(log.Fields{
			"userID": foundUser.ID,
		}).Info("user login")

		return c.JSON(fiber.Map{"token": t})
	}
}
