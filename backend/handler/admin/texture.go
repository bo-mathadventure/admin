package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
	"time"
)

func NewAdminTextureHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminTexture(ctx, db))
	app.Post("/", postAdminTexture(ctx, db))
	app.Get("/:id", getAdminTextureID(ctx, db))
	app.Put("/:id", putAdminTextureID(ctx, db))
	app.Delete("/:id", deleteAdminTextureID(ctx, db))
}

type AdminTextureResponse struct {
	ID        int       `json:"id"`
	Texture   string    `json:"texture" example:"eyes1"`
	Layer     string    `json:"layer" enums:"woka,body,hair,eyes,hat,accessory,clothes,companion"`
	URL       string    `json:"url" example:"%FRONTEND_URL%/public/resources/customisation/character_eyes/character_eyes1.png"`
	Tags      []string  `json:"tags" example:"editor"`
	CreatedAt time.Time `json:"createdAt" example:"2006-01-02T15:04:05Z07:00"`
}

// getAdminTexture godoc
//
//	@Summary		List textures
//	@Description	Get all textures. Requires permission admin.texture.view or admin.texture.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		AdminTextureResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/texture [get]
func getAdminTexture(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_TEXTURE_VIEW || utils.PERMISSION_TEXTURE_EDIT
		return handler.HandleSuccess(c)
	}
}

type CreateTexture struct {
	Texture string   `json:"texture" example:"eyes1"`
	Layer   string   `json:"layer" enums:"woka,body,hair,eyes,hat,accessory,clothes,companion"`
	URL     string   `json:"url" example:"%FRONTEND_URL%/public/resources/customisation/character_eyes/character_eyes1.png"`
	Tags    []string `json:"tags" example:"editor"`
}

// postAdminTexture godoc
//
//	@Summary		Create texture
//	@Description	Requires permission admin.texture.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		CreateTexture	true	"-"
//	@Success		200		{object}	AdminTextureResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/texture [post]
func postAdminTexture(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_TEXTURE_EDIT
		return handler.HandleSuccess(c)
	}
}

// getAdminTextureID godoc
//
//	@Summary		Get specific texture
//	@Description	Get texture by ID. Requires permission admin.texture.view or admin.texture.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Texture ID"
//	@Success		200	{object}	AdminTextureResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/texture/{id} [get]
func getAdminTextureID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_TEXTURE_VIEW || utils.PERMISSION_TEXTURE_EDIT
		return handler.HandleSuccess(c)
	}
}

type UpdateTexture struct {
	Texture string   `json:"texture" example:"eyes1"`
	Layer   string   `json:"layer" enums:"woka,body,hair,eyes,hat,accessory,clothes,companion"`
	URL     string   `json:"url" example:"%FRONTEND_URL%/public/resources/customisation/character_eyes/character_eyes1.png"`
	Tags    []string `json:"tags" example:"editor"`
}

// putAdminTextureID godoc
//
//	@Summary		Update texture
//	@Description	Update texture by ID. Requires permission admin.texture.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		UpdateTexture	true	"-"
//	@Param			id		path		int				true	"Texture ID"
//	@Success		200		{object}	AdminTextureResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/texture/{id} [put]
func putAdminTextureID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_TEXTURE_EDIT
		return handler.HandleSuccess(c)
	}
}

// deleteAdminTextureID godoc
//
//	@Summary		Delete texture
//	@Description	Delete texture by ID. Requires permission admin.texture.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Texture ID"
//	@Success		200	{object}	handler.APIResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/texture/{id} [delete]
func deleteAdminTextureID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_TEXTURE_EDIT
		return handler.HandleSuccess(c)
	}
}
