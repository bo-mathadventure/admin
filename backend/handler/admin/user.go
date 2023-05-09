package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

// NewAdminUserHandler initialize routes for the given router
func NewAdminUserHandler(ctx context.Context, app fiber.Router, db *ent.Client) {
	app.Get("/", getAdminUser(ctx, db))
	app.Post("/invite", postAdminUser(ctx, db))
	app.Get("/:id", getAdminUserID(ctx, db))
	app.Put("/:id", putAdminUserID(ctx, db))
	app.Delete("/:id", deleteAdminUserID(ctx, db))
}

type adminUserResponse struct {
	ID              int                   `json:"id"`
	UUID            string                `json:"uuid" example:"60948703-fca9-4491-b3bc-588188d93eb3"`
	Email           string                `json:"email" example:"bob@example.com"`
	Username        string                `json:"username" example:"Bob"`
	SsoIdentifier   string                `json:"ssoIdentifier"`
	Permissions     []string              `json:"permissions" example:"admin.editor,admin.user.edit"`
	UserPermissions []string              `json:"userPermissions" example:"admin.editor,admin.user.edit"`
	Tags            []string              `json:"tags" example:"admin,mod,editor,student"`
	UserTags        []string              `json:"userTags" example:"admin,mod,editor,student"`
	LastLogin       string                `json:"lastLogin" example:"2006-01-02T15:04:05Z07:00" validate:"omitempty"`
	CreatedAt       string                `json:"createdAt" example:"2006-01-02T15:04:05Z07:00"`
	Groups          []*adminGroupResponse `json:"groups"`
}

func responseAdminUserResponse(this *ent.User) *adminUserResponse {
	return &adminUserResponse{
		ID:              this.ID,
		UUID:            this.UUID,
		Email:           this.Email,
		Username:        this.Username,
		SsoIdentifier:   this.SsoIdentifier,
		Permissions:     utils.CombinePermissions(this),
		UserPermissions: this.Permissions,
		Tags:            utils.CombineTags(this),
		UserTags:        this.Tags,
		Groups:          responseAdminGroupResponses(this.Edges.Groups),
		LastLogin:       this.LastLogin.Format(time.RFC3339),
		CreatedAt:       this.CreatedAt.Format(time.RFC3339),
	}
}

func responseAdminUserResponses(this []*ent.User) []*adminUserResponse {
	data := make([]*adminUserResponse, len(this))
	for i, e := range this {
		data[i] = responseAdminUserResponse(e)
	}
	return data
}

// getAdminUser godoc
//
//	@Summary		List users
//	@Description	Get all users. Requires permission admin.user.view or admin.user.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		adminUserResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/user [get]
func getAdminUser(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userID)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PermissionUserView, utils.PermissionUserEdit}) {
			return handler.HandleInvalidPermissions(c)
		}

		allEntires, err := db.User.Query().All(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminUserResponses(allEntires))
	}
}

type createUser struct {
	Email    string `json:"email" example:"bob@example.com" validate:"required,email"`
	Username string `json:"username" example:"Bob" validate:"required,alphaunicode,min=3,max=16"`
	Password string `json:"password" example:"my$ecur3P4$$word" validate:"required,min=8"`
}

// postAdminUser godoc
//
//	@Summary		Invite user
//	@Description	Invite/Create new user. Requires permission admin.user.invite
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		createUser	true	"-"
//	@Success		200		{object}	adminUserResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/user/invite [post]
func postAdminUser(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userID)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PermissionUserInvite}) {
			return handler.HandleInvalidPermissions(c)
		}

		req := new(createUser)
		if err := c.BodyParser(req); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if valid, err := handler.ValidateStruct(c, req); !valid {
			return err
		}

		hashPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		newEntry, err := db.User.Create().SetEmail(req.Email).SetUsername(req.Username).SetPassword(hashPassword).Save(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminUserResponse(newEntry))
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
//	@Success		200	{object}	adminUserResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/user/{id} [get]
func getAdminUserID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userID)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PermissionUserView, utils.PermissionUserEdit}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		foundEntry, err := db.User.Query().WithGroups().Where(user.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminUserResponse(foundEntry))
	}
}

type updateUser struct {
	Email       string   `json:"email" example:"bob@example.com" validate:"required,email"`
	Username    string   `json:"username" example:"Bob" validate:"required,min=3,max=16"`
	Password    string   `json:"password" validate:"omitempty,min=8"`
	Permissions []string `json:"permissions" example:"admin.editor,admin.user.edit" validate:"required"`
	Tags        []string `json:"tags" example:"admin,mod,editor,student" validate:"required"`
	Groups      []int    `json:"groups" example:"10" validate:"required"`
}

// putAdminUserID godoc
//
//	@Summary		Update user
//	@Description	Update user by ID. Requires permission admin.user.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		updateUser	true	"-"
//	@Param			id		path		int			true	"User ID"
//	@Success		200		{object}	adminUserResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/user/{id} [put]
func putAdminUserID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userID)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PermissionUserEdit}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		req := new(updateUser)
		if err := c.BodyParser(req); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if valid, err := handler.ValidateStruct(c, req); !valid {
			return err
		}

		found, err := db.User.Query().WithGroups().Where(user.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		foundUpdate := found.Update().SetEmail(req.Email).SetUsername(req.Username).SetPermissions(req.Permissions).AddGroupIDs(req.Groups...).SetTags(req.Tags)
		if req.Password != "" {
			hashPassword, err := utils.HashPassword(req.Password)
			if err != nil {
				return handler.HandleInternalError(c, err)
			}
			foundUpdate = foundUpdate.SetPassword(hashPassword)
		}

		newFound, err := foundUpdate.Save(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminUserResponse(newFound))
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
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userID)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PermissionUserEdit}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		err = db.User.DeleteOneID(pathID).Exec(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return handler.HandleSuccess(c)
	}
}
