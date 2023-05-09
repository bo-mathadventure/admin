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

type registerRequest struct {
	EMail                    string `json:"email" example:"bob@example.com" validate:"required,email"`
	Username                 string `json:"username" example:"Bob" validate:"required,alphaunicode,min=3,max=16"`
	Language                 string `json:"language" example:"de" validate:"omitempty"`
	ClearTextPassword        string `json:"password" example:"my$ecur3P4$$word" validate:"required,min=8"`
	ClearTextPasswordConfirm string `json:"confirmPassword" example:"my$ecur3P4$$word" validate:"required,min=8,eqcsfield=ClearTextPassword"`
}

// Register godoc
//
//	@Summary		Register new user
//	@Description	Start a registration of a new user. Only works when registration are enabled
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			params	body		registerRequest	true	"-"
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

		req := new(registerRequest)
		if err := c.BodyParser(req); err != nil {
			return HandleBodyParseError(c, err)
		}

		if valid, err := ValidateStruct(c, req); !valid {
			return err
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
