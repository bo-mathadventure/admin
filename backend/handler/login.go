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
	EMail             string `json:"email" example:"bob@example.com" validate:"required,email"`
	ClearTextPassword string `json:"password" example:"my$ecur3P4$$word" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Login godoc
//
//	@Summary		Get Login Token
//	@Description	Do a login with user credentials (email/password) when a password is set
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			params	body		LoginRequest	true	"-"
//	@Success		200		{object}	LoginResponse
//	@Failure		400		{object}	APIResponse
//	@Failure		404		{object}	APIResponse
//	@Failure		500		{object}	APIResponse
//	@Router			/auth/login [post]
func Login(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(LoginRequest)
		if err := c.BodyParser(req); err != nil {
			return HandleBodyParseError(c, err)
		}

		if valid, err := ValidateStruct(c, req); !valid {
			return err
		}

		foundUser, err := db.User.Query().Where(user.Email(email.Normalize(req.EMail))).First(ctx)
		if err != nil || foundUser == nil || !utils.CheckPasswordHash(req.ClearTextPassword, foundUser.Password) {
			return HandleInvalidLogin(c)
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":            foundUser.ID,
			"email":         foundUser.Email,
			"ssoIdentifier": nil,
			"exp":           time.Now().Add(time.Hour * 72).Unix(),
		})

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(config.GetConfig().WorkadventureSecretKey))
		if err != nil {
			return HandleInternalError(c, err)
		}

		_, err = foundUser.Update().SetLastLogin(time.Now()).Save(ctx)
		if err != nil {
			return HandleInternalError(c, err)
		}

		log.WithFields(log.Fields{
			"userID": foundUser.ID,
		}).Info("user login")

		return c.JSON(LoginResponse{t})
	}
}
