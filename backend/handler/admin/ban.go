package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func NewAdminBanHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminBan(ctx, db))
	app.Post("/", postAdminBan(ctx, db))
	app.Get("/:id", getAdminBanID(ctx, db))
	app.Delete("/:id", deleteAdminBanID(ctx, db))
}

type AdminBanResponse struct {
	ID         int       `json:"id"`
	Check      string    `json:"check" example:"60948703-fca9-4491-b3bc-588188d93eb3"`
	Message    string    `json:"message" example:"banned for verbal abuse"`
	ValidUntil time.Time `json:"validUntil" example:"2006-01-02T15:04:05Z07:00" validate:"optional"`
	CreatedAt  time.Time `json:"createdAt" example:"2006-01-02T15:04:05Z07:00"`
}

// getAdminBan godoc
//
//	@Summary		List bans
//	@Description	Get all bans. Requires permission admin.ban.view or admin.ban.create or admin.ban.delete
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		AdminBanResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/ban [get]
func getAdminBan(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_BAN_VIEW, utils.PERMISSION_BAN_CREATE, utils.PERMISSION_BAN_DELETE}) {
			return handler.HandleInvalidPermissions(c)
		}
		return handler.HandleSuccess(c)
	}
}

type CreateBan struct {
	Check      string    `json:"check" example:"60948703-fca9-4491-b3bc-588188d93eb3"`
	Message    string    `json:"message" example:"banned for verbal abuse"`
	ValidUntil time.Time `json:"validUntil" example:"2006-01-02T15:04:05Z07:00" validate:"optional"`
}

// postAdminBan godoc
//
//	@Summary		Create ban
//	@Description	Create a ban for an identifier or ip address. Requires permission admin.ban.create
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		CreateBan	true	"-"
//	@Success		200		{object}	AdminBanResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/ban [post]
func postAdminBan(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_BAN_CREATE}) {
			return handler.HandleInvalidPermissions(c)
		}
		return handler.HandleSuccess(c)
	}
}

// getAdminBanID godoc
//
//	@Summary		Get specific ban
//	@Description	Get ban by ID. Requires permission admin.ban.view or admin.ban.create or admin.ban.delete
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Ban ID"
//	@Success		200	{object}	AdminBanResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/ban/{id} [get]
func getAdminBanID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_BAN_VIEW, utils.PERMISSION_BAN_CREATE, utils.PERMISSION_BAN_DELETE}) {
			return handler.HandleInvalidPermissions(c)
		}
		return handler.HandleSuccess(c)
	}
}

// deleteAdminBanID godoc
//
//	@Summary		Delete ban
//	@Description	Delete ban by ID. Sets validUntil to now() to keep history. Requires permission admin.ban.delete
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Ban ID"
//	@Success		200	{object}	handler.APIResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/ban/{id} [delete]
func deleteAdminBanID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_BAN_DELETE}) {
			return handler.HandleInvalidPermissions(c)
		}
		return handler.HandleSuccess(c)
	}
}
