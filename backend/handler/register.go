package handler

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/utils"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type RegisterRequest struct {
	EMail                    string `json:"email"`
	Username                 string `json:"username"`
	Language                 string `json:"language"`
	ClearTextPassword        string `json:"password"`
	ClearTextPasswordConfirm string `json:"confirmPassword"`
}

func Register(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(RegisterRequest)
		if err := c.BodyParser(req); err != nil {
			return HandleBodyParseError(c, err)
		}

		if req.EMail == "" || req.ClearTextPassword == "" || req.Username == "" {
			return HandleError(c, "ERR_INVALID_REQUEST")
		}

		err := email.Validate(req.EMail)
		if err != nil {
			return HandleError(c, "ERR_EMAIL_INVALID")
		}

		if req.ClearTextPassword != req.ClearTextPasswordConfirm {
			return HandleError(c, "ERR_PASSWORD_EQUAL")
		}

		hashedPassword, err := utils.HashPassword(req.ClearTextPassword)
		if err != nil {
			return HandleInternalError(c, err)
		}

		existingUser, err := db.User.Query().Where(user.Email(email.Normalize(req.EMail))).Count(ctx)
		if existingUser > 0 || err != nil {
			return HandleError(c, "ERR_USER_EXISTS")
		}

		newUser, err := db.User.Create().SetEmail(email.Normalize(req.EMail)).SetPassword(hashedPassword).SetUsername(req.Username).Save(ctx)
		if err != nil || newUser == nil {
			return HandleInternalError(c, err)
		}
		log.WithFields(log.Fields{
			"userID": newUser.ID,
		}).Info("user registered")

		return HandleSuccess(c)
	}
}
