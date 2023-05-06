package handler

import (
	"context"
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/utils"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type RegisterRequest struct {
	EMail                    string `json:"email" example:"bob@example.com"`
	Username                 string `json:"username" example:"Bob"`
	Language                 string `json:"language" example:"de"`
	ClearTextPassword        string `json:"password" example:"my$ecur3P4$$word"`
	ClearTextPasswordConfirm string `json:"confirmPassword" example:"my$ecur3P4$$word"`
}

// Register godoc
//
//	@Summary		Register new user
//	@Description	Start a registration of a new user. Only works when registration are enabled
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			params	body		RegisterRequest	true	"-"
//	@Success		200		{object}	APIResponse
//	@Failure		400		{object}	APIResponse
//	@Failure		404		{object}	APIResponse
//	@Failure		500		{object}	APIResponse
//	@Router			/auth/register [post]
func Register(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !config.GetConfig().EnableRegistration {
			return HandleError(c, "ERR_REGISTRATION_DISABLED")
		}

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
