package handler

import (
	"context"
	"github.com/bo-mathadventure/admin/encoder"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/utils"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func NewUserHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getMe(ctx, db))
	app.Put("/", updateUser(ctx, db))
}

func getMe(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return HandleInternalError(c, err)
		}

		result := encoder.ParseUser(thisUser)
		return c.JSON(&result)
	}
}

type UpdateUserRequest struct {
	EMail                    string `json:"email"`
	ClearTextPassword        string `json:"newPassword"`
	ClearTextPasswordConfirm string `json:"confirmPassword"`
	ClearTextCurrentPassword string `json:"password"`
}

func updateUser(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		req := new(UpdateUserRequest)
		if err := c.BodyParser(req); err != nil {
			return HandleBodyParseError(c, err)
		}

		foundUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)

		if req.EMail != "" || req.ClearTextPassword != "" {
			if err != nil || foundUser == nil || !utils.CheckPasswordHash(req.ClearTextCurrentPassword, foundUser.Password) {
				return HandleError(c, "ERR_INVALID_PASSWORD")
			}
		}

		update := foundUser.Update()

		if req.EMail != "" {
			err := email.Validate(req.EMail)
			if err != nil {
				return HandleError(c, "ERR_EMAIL_INVALID")
			}

			existingUser, err := db.User.Query().Where(user.Email(email.Normalize(req.EMail))).Where(user.IDNotIn(userId)).Count(ctx)
			if existingUser > 0 || err != nil {
				return HandleError(c, "ERR_USER_EXISTS")
			}

			update = update.SetEmail(email.Normalize(req.EMail))
		}

		if req.ClearTextPassword != "" && req.ClearTextPasswordConfirm == "" || req.ClearTextPassword == "" && req.ClearTextPasswordConfirm != "" {
			return HandleError(c, "ERR_INVALID_REQUEST")
		}
		if req.ClearTextPassword != "" && req.ClearTextPasswordConfirm != "" {
			if req.ClearTextPassword != req.ClearTextPasswordConfirm {
				return HandleError(c, "ERR_PASSWORD_EQUAL")
			}

			hashedPassword, err := utils.HashPassword(req.ClearTextPassword)
			if err != nil {
				return HandleInternalError(c, err)
			}
			update = update.SetPassword(hashedPassword)
		}

		_, err = update.Save(ctx)
		if err != nil {
			return HandleInternalError(c, err)
		}

		return HandleSuccess(c)
	}
}
