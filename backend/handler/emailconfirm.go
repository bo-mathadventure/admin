package handler

import (
	"context"
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/mailer"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
)

// ResendConfirmEmail godoc
//
//	@Summary		Resend mail confirmation
//	@Description	Resend mail confirmation if user not confirmed
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			params	body		updateUserRequest	true	"-"
//	@Success		200		{object}	APIResponse
//	@Failure		400		{object}	APIResponse
//	@Failure		404		{object}	APIResponse
//	@Failure		500		{object}	APIResponse
//	@Router			/auth/resendConfirmation [post]
func ResendConfirmEmail(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !config.GetConfig().RegistrationEMailConfirmation {
			return HandleSuccess(c)
		}

		req := new(updateUserRequest)
		if err := c.BodyParser(req); err != nil {
			return HandleBodyParseError(c, err)
		}

		if req.EMail == "" {
			return HandleSuccess(c)
		}

		foundUser, err := db.User.Query().Where(user.Email(email.Normalize(req.EMail))).First(ctx)
		if err != nil || foundUser.EmailConfirmed {
			return HandleSuccess(c)
		}

		_, err = db.Token.Create().SetUser(foundUser).SetAction(mailer.ActionConfirmEmail).Save(ctx)
		if err != nil {
			return HandleInternalError(c, err)
		}

		return HandleSuccess(c)
	}
}
