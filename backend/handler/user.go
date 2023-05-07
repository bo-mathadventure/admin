package handler

import (
	"context"
	"fmt"
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/utils"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func NewUserHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getMe(ctx, db))
	app.Put("/", updateUser(ctx, db))
	app.Get("/token", getTokenLogin(ctx, db))
}

type UserResponse struct {
	UUID        string        `json:"uuid"`
	Email       string        `json:"email"`
	Username    string        `json:"username"`
	Permissions []string      `json:"permissions"`
	Tags        []string      `json:"tags"`
	LastLogin   time.Time     `json:"lastLogin" validate:"omitempty"`
	CreatedAt   time.Time     `json:"createdAt"`
	Config      config.Config `json:"config"`
}

func responseUserResponse(thisUser *ent.User) *UserResponse {
	return &UserResponse{
		UUID:        thisUser.UUID,
		Email:       thisUser.Email,
		Username:    thisUser.Username,
		Permissions: utils.CombinePermissions(thisUser),
		Tags:        utils.CombineTags(thisUser),
		LastLogin:   thisUser.LastLogin,
		CreatedAt:   thisUser.CreatedAt,
		Config:      config.GetConfig(),
	}
}

// getMe godoc
//
//	@Summary		User Info
//	@Description	Get user details of logged-in user
//	@Security		ApiKeyAuth
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	UserResponse
//	@Failure		401	{object}	APIResponse
//	@Failure		500	{object}	APIResponse
//	@Router			/system/user [get]
func getMe(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return HandleInternalError(c, err)
		}

		return c.JSON(responseUserResponse(thisUser))
	}
}

type UpdateUserRequest struct {
	EMail                    string `json:"email" example:"bob@exameple.com" format:"email" validate:"omitempty" validate:"required,email"`
	ClearTextPassword        string `json:"newPassword" example:"my$ecur3P4$$word" validate:"omitempty" validate:"omitempty,min=8"`
	ClearTextPasswordConfirm string `json:"confirmPassword" example:"my$ecur3P4$$word" validate:"omitempty" validate:"omitempty,min=8,eqcsfield=ClearTextPassword"`
	ClearTextCurrentPassword string `json:"password" example:"my$ecur3P4$$word" validate:"required" validate:"required"`
}

// updateUser godoc
//
//	@Summary		Update User
//	@Description	Update details of the logged-in user
//	@Security		ApiKeyAuth
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			params	body		UpdateUserRequest	true	"-"
//	@Success		200		{object}	UserResponse
//	@Failure		400		{object}	APIResponse
//	@Failure		401		{object}	APIResponse
//	@Failure		404		{object}	APIResponse
//	@Failure		500		{object}	APIResponse
//	@Router			/system/user [put]
func updateUser(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		req := new(UpdateUserRequest)
		if err := c.BodyParser(req); err != nil {
			return HandleBodyParseError(c, err)
		}

		if valid, err := ValidateStruct(c, req); !valid {
			return err
		}

		foundUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)

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

		updatedUser, err := update.Save(ctx)
		if err != nil {
			return HandleInternalError(c, err)
		}

		return c.JSON(responseUserResponse(updatedUser))
	}
}

// getTokenLogin godoc
//
//	@Summary		Workadventure token
//	@Description	Generate JWT Token of logged-in user and directly redirect user to Workadventure
//	@Security		ApiKeyAuth
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		302	{object}	nil
//	@Failure		401	{object}	APIResponse
//	@Failure		500	{object}	APIResponse
//	@Router			/system/user/token [get]
func getTokenLogin(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return HandleInternalError(c, err)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"identifier":  thisUser.UUID, // fixme may be also email
			"accessToken": nil,
			"username":    thisUser.Username,
		})

		tokenString, err := token.SignedString([]byte(config.GetConfig().WorkadventureSecretKey))
		if err != nil {
			return HandleInternalError(c, err)
		}
		return c.Redirect(fmt.Sprintf("%s?token=%s", config.GetConfig().WorkadventureURL, tokenString))
	}
}
