package admin

import (
	"context"
	"fmt"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/textures"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strconv"
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
	ID        int      `json:"id"`
	Texture   string   `json:"texture" example:"eyes1"`
	Layer     string   `json:"layer" enums:"woka,body,hair,eyes,hat,accessory,clothes,companion"`
	URL       string   `json:"url" example:"%FRONTEND_URL%/public/resources/customisation/character_eyes/character_eyes1.png"`
	Tags      []string `json:"tags" example:"editor"`
	CreatedAt string   `json:"createdAt" example:"2006-01-02T15:04:05Z07:00"`
}

func responseAdminTextureResponse(this *ent.Textures) *AdminTextureResponse {
	return &AdminTextureResponse{
		ID:        this.ID,
		Texture:   this.Texture,
		Layer:     this.Layer,
		URL:       this.URL,
		Tags:      this.Tags,
		CreatedAt: this.CreatedAt.Format(time.RFC3339),
	}
}

func responseAdminTextureResponses(this []*ent.Textures) []*AdminTextureResponse {
	data := make([]*AdminTextureResponse, len(this))
	for i, e := range this {
		data[i] = responseAdminTextureResponse(e)
	}
	return data
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
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_TEXTURE_VIEW, utils.PERMISSION_TEXTURE_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		allEntires, err := db.Textures.Query().All(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminTextureResponses(allEntires))
	}
}

type CreateTexture struct {
	Texture string   `json:"texture" form:"texture" example:"eyes1" validate:"required"`
	Layer   string   `json:"layer" form:"layer" enums:"woka,body,hair,eyes,hat,accessory,clothes,companion" validate:"required"`
	Tags    []string `json:"tags" json:"tags" example:"editor" validate:"required"`
}

// postAdminTexture godoc
//
//	@Summary		Create texture
//	@Description	Upload file via resource field. Requires permission admin.texture.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			mpfd
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
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_TEXTURE_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		req := new(CreateTexture)
		if err := c.BodyParser(req); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if valid, err := handler.ValidateStruct(c, req); !valid {
			return err
		}

		form, err := c.MultipartForm()
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		files := form.File["resource"]
		if len(files) != 1 {
			return handler.HandleErrorCode(c, fiber.StatusBadRequest, "ERR_RESOURCE")
		}

		file := files[0]
		newFilename := fmt.Sprintf("%s.%s", uuid.NewString(), filepath.Ext(file.Filename))
		err = os.MkdirAll("./public/upload/", 0660)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}
		err = c.SaveFile(file, fmt.Sprintf("./public/upload/%s", newFilename))
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		newEntry, err := db.Textures.Create().
			SetTexture(req.Texture).
			SetLayer(req.Layer).
			SetTags(req.Tags).
			SetURL("%FRONTEND_URL%/public/upload/" + newFilename).
			Save(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminTextureResponse(newEntry))
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
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_TEXTURE_VIEW, utils.PERMISSION_TEXTURE_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		foundEntry, err := db.Textures.Query().Where(textures.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminTextureResponse(foundEntry))
	}
}

type UpdateTexture struct {
	Texture string   `json:"texture" form:"texture" example:"eyes1" validate:"required"`
	Layer   string   `json:"layer" form:"layer" enums:"woka,body,hair,eyes,hat,accessory,clothes,companion" validate:"required"`
	Tags    []string `json:"tags" json:"tags" example:"editor" validate:"required"`
}

// putAdminTextureID godoc
//
//	@Summary		Update texture
//	@Description	Update texture by ID. Requires permission admin.texture.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			mpfd
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
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_TEXTURE_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		req := new(UpdateTexture)
		if err := c.BodyParser(req); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if valid, err := handler.ValidateStruct(c, req); !valid {
			return err
		}

		found, err := db.Textures.Query().Where(textures.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		foundUpdate := found.Update().
			SetTexture(req.Texture).
			SetLayer(req.Layer).
			SetTags(req.Tags)

		form, err := c.MultipartForm()
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		files := form.File["resource"]
		if len(files) == 1 {
			file := files[0]
			newFilename := fmt.Sprintf("%s.%s", uuid.NewString(), filepath.Ext(file.Filename))
			err = os.MkdirAll("./public/upload/", 0660)
			if err != nil {
				return handler.HandleInternalError(c, err)
			}
			err = c.SaveFile(file, fmt.Sprintf("./public/upload/%s", newFilename))
			if err != nil {
				return handler.HandleInternalError(c, err)
			}

			foundUpdate = foundUpdate.SetURL("%FRONTEND_URL%/public/upload/" + newFilename)
		}

		newFound, err := foundUpdate.Save(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminTextureResponse(newFound))
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
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_TEXTURE_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		err = db.Textures.DeleteOneID(pathID).Exec(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return handler.HandleSuccess(c)
	}
}
