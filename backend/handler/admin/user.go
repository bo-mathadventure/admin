package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
	"time"
)

func NewAdminUserHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminUser(ctx, db))
	app.Post("/invite", postAdminUser(ctx, db))
	app.Get("/:id", getAdminUserID(ctx, db))
	app.Put("/:id", putAdminUserID(ctx, db))
	app.Delete("/:id", deleteAdminUserID(ctx, db))
}

type AdminUserResponse struct {
	ID            int       `json:"id"`
	UUID          string    `json:"uuid" example:"60948703-fca9-4491-b3bc-588188d93eb3"`
	Email         string    `json:"email" example:"bob@example.com"`
	Username      string    `json:"username" example:"Bob"`
	SsoIdentifier string    `json:"ssoIdentifier"`
	VCardURL      string    `json:"vCardURL" validate:"optional"`
	Permissions   []string  `json:"permissions" example:"admin.editor,admin.user.*"`
	Tags          []string  `json:"tags" example:"admin,mod,editor,student"`
	LastLogin     time.Time `json:"lastLogin" example:"2006-01-02T15:04:05Z07:00" validate:"optional"`
	CreatedAt     time.Time `json:"createdAt" example:"2006-01-02T15:04:05Z07:00"`
}

// getAdminUser godoc
//
//	@Summary		List users
//	@Description	Get all users. Requires permission admin.user.view or admin.user.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		AdminUserResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/user [get]
func getAdminUser(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_USER_VIEW || utils.PERMISSION_USER_EDIT
		return handler.HandleSuccess(c)
	}
}

type CreateUser struct {
	Email    string `json:"email" example:"bob@example.com"`
	Username string `json:"username" example:"Bob"`
	Password string `json:"password" example:"my$ecur3P4$$word"`
}

// postAdminUser godoc
//
//	@Summary		Invite user
//	@Description	Invite/Create new user. Requires permission admin.user.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		CreateUser	true	"-"
//	@Success		200		{object}	AdminUserResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/user/invite [post]
func postAdminUser(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_USER_EDIT
		return handler.HandleSuccess(c)
	}
}

// getAdminUserID godoc
//
//	@Summary		Get specific user
//	@Description	Get user by ID. Requires permission admin.user.view or admin.user.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	AdminUserResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/user/{id} [get]
func getAdminUserID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_USER_VIEW || utils.PERMISSION_USER_EDIT
		return handler.HandleSuccess(c)
	}
}

type UpdateUser struct {
	Email       string   `json:"email" example:"bob@example.com"`
	Username    string   `json:"username" example:"Bob"`
	VCardURL    string   `json:"vCardURL" validate:"optional"`
	Permissions []string `json:"permissions" example:"admin.editor,admin.user.*"`
	Tags        []string `json:"tags" example:"admin,mod,editor,student"`
}

// putAdminUserID godoc
//
//	@Summary		Update user
//	@Description	Update user by ID. Requires permission admin.user.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		UpdateUser	true	"-"
//	@Param			id		path		int			true	"User ID"
//	@Success		200		{object}	AdminUserResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/user/{id} [put]
func putAdminUserID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_USER_EDIT
		return handler.HandleSuccess(c)
	}
}

// deleteAdminUserID godoc
//
//	@Summary		Delete user
//	@Description	Delete user by ID. Requires permission admin.user.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	handler.APIResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/user/{id} [delete]
func deleteAdminUserID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_USER_EDIT
		return handler.HandleSuccess(c)
	}
}
