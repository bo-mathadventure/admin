package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/group"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

func NewAdminGroupHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminGroup(ctx, db))
	app.Post("/", postAdminGroup(ctx, db))
	app.Get("/:id", getAdminGroupID(ctx, db))
	app.Put("/:id", putAdminGroupID(ctx, db))
	app.Delete("/:id", deleteAdminGroupID(ctx, db))
}

type AdminGroupResponse struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty" example:"administration"`
	Description string   `json:"description,omitempty" example:""`
	Permissions []string `json:"permissions,omitempty" example:"admin.user.view,admin.user.view"`
	Tags        []string `json:"tags,omitempty" example:"admin"`
	Deleteable  bool     `json:"deleteable,omitempty" `
	CreatedAt   string   `json:"createdAt" example:"2006-01-02T15:04:05Z07:00"`
}

func responseAdminGroupResponse(this *ent.Group) *AdminGroupResponse {
	return &AdminGroupResponse{
		ID:          this.ID,
		Name:        this.Name,
		Description: this.Description,
		Permissions: this.Permissions,
		Tags:        this.Tags,
		Deleteable:  this.Edges.Users == nil || len(this.Edges.Users) == 0,
		CreatedAt:   this.CreatedAt.Format(time.RFC3339),
	}
}

func responseAdminGroupResponses(this []*ent.Group) []*AdminGroupResponse {
	data := make([]*AdminGroupResponse, len(this))
	for i, e := range this {
		data[i] = responseAdminGroupResponse(e)
	}
	return data
}

// getAdminGroup godoc
//
//	@Summary		List Groups
//	@Description	Get all groups. Requires permission admin.group.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		AdminGroupResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/group [get]
func getAdminGroup(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_GROUP_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		allEntires, err := db.Group.Query().WithUsers().All(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminGroupResponses(allEntires))
	}
}

type CreateGroup struct {
	Name string `json:"name,omitempty" example:"administration" validate:"required,min=3,max=32"`
}

// postAdminGroup godoc
//
//	@Summary		Create Group
//	@Description	Create new Group. Requires permission admin.group.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		CreateGroup	true	"-"
//	@Success		200		{object}	AdminGroupResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/group [post]
func postAdminGroup(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_GROUP_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		req := new(CreateGroup)
		if err := c.BodyParser(req); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if valid, err := handler.ValidateStruct(c, req); !valid {
			return err
		}

		newEntry, err := db.Group.Create().SetName(req.Name).SetDescription("").Save(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminGroupResponse(newEntry))
	}
}

// getAdminGroupID godoc
//
//	@Summary		Get specific group
//	@Description	Get group by ID. Requires permission admin.group.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Group ID"
//	@Success		200	{object}	AdminGroupResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/group/{id} [get]
func getAdminGroupID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_GROUP_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		foundEntry, err := db.Group.Query().WithUsers().Where(group.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminGroupResponse(foundEntry))
	}
}

type UpdateGroup struct {
	Name        string   `json:"name,omitempty" example:"administration" validate:"required,min=3,max=32"`
	Description string   `json:"description,omitempty" example:""`
	Permissions []string `json:"permissions,omitempty" example:"admin.user.view,admin.user.view"`
	Tags        []string `json:"tags,omitempty" example:"admin"`
}

// putAdminGroupID godoc
//
//	@Summary		Update group
//	@Description	Update group by ID. Requires permission admin.group.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		UpdateGroup	true	"-"
//	@Param			id		path		int			true	"Group ID"
//	@Success		200		{object}	AdminGroupResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/group/{id} [put]
func putAdminGroupID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_USER_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		req := new(UpdateGroup)
		if err := c.BodyParser(req); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if valid, err := handler.ValidateStruct(c, req); !valid {
			return err
		}

		found, err := db.Group.Query().Where(group.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		newFound, err := found.Update().SetName(req.Name).SetDescription(req.Description).SetTags(req.Tags).SetPermissions(req.Permissions).Save(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminGroupResponse(newFound))
	}
}

// deleteAdminGroupID godoc
//
//	@Summary		Delete group
//	@Description	Delete group by ID. Requires permission admin.group.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Group ID"
//	@Success		200	{object}	handler.APIResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/group/{id} [delete]
func deleteAdminGroupID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_GROUP_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		err = db.Group.DeleteOneID(pathID).Exec(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return handler.HandleSuccess(c)
	}
}
