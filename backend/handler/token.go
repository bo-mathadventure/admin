package handler

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/token"
	"github.com/bo-mathadventure/admin/mailer"
	"github.com/bo-mathadventure/admin/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

// Token godoc
//
//	@Summary		Token
//	@Description	execute token actions
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			token	query		string				true	"-"
//	@Param			params	body		updateUserRequest	true	"-"
//	@Success		200		{object}	userResponse
//	@Failure		400		{object}	APIResponse
//	@Failure		401		{object}	APIResponse
//	@Failure		404		{object}	APIResponse
//	@Failure		500		{object}	APIResponse
//	@Router			/auth/token [post]
func Token(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenValue := c.Query("token", "")
		if tokenValue == "" {
			return HandleSuccess(c)
		}

		queryToken, err := db.Token.Query().
			Where(token.Token(tokenValue)).
			Where(token.ValidUntilGTE(time.Now())).
			WithUser().
			Only(ctx)
		if err != nil {
			return HandleError(c, "INVALID_TOKEN")
		}

		if queryToken.Action == mailer.ActionConfirmEmail {
			_, err = queryToken.Edges.User.Update().SetEmailConfirmed(true).Save(ctx)
			if err != nil {
				return HandleInternalError(c, err)
			}
			_ = db.Token.DeleteOne(queryToken).Exec(ctx)
		} else if queryToken.Action == mailer.ActionPasswordReset {
			req := new(updateUserRequest)
			if err := c.BodyParser(req); err != nil {
				return HandleBodyParseError(c, err)
			}

			if req.ClearTextPassword == "" {
				return HandleError(c, "ERR_INVALID_REQUEST")
			}

			hashedPassword, err := utils.HashPassword(req.ClearTextPassword)
			if err != nil {
				return HandleInternalError(c, err)
			}

			_, err = queryToken.Edges.User.Update().SetPassword(hashedPassword).Save(ctx)
			if err != nil {
				return HandleInternalError(c, err)
			}
		}

		_ = db.Token.DeleteOne(queryToken).Exec(ctx)
		return HandleSuccess(c)
	}
}
