package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/announcement"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func NewAdminAnnouncementHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	err := handler.Validate.RegisterValidation("announcement_type_validation", func(fl validator.FieldLevel) bool {
		return utils.Contains([]string{"ban", "warning"}, fl.Field().String())
	})
	if err != nil {
		log.WithError(err).WithField("validation", "policy_validation").Panic("failed to setup custom validation")
	}

	app.Get("/", getAdminAnnouncement(ctx, db))
	app.Post("/", postAdminAnnouncement(ctx, db))
	app.Get("/:id", getAdminAnnouncementID(ctx, db))
	app.Put("/:id", putAdminAnnouncementID(ctx, db))
	app.Delete("/:id", deleteAdminAnnouncementID(ctx, db))
}

type AdminAnnouncementResponse struct {
	ID         int    `json:"id"`
	Type       string `json:"type" enums:"ban,warning"`
	Message    string `json:"message" example:"This is an example alert"`
	CreatedAt  string `json:"createdAt" example:"2006-01-02T15:04:05Z07:00"`
	ValidUntil string `json:"validUntil" example:"2006-01-02T15:04:05Z07:00" validate:"omitempty"`
}

func responseAdminAnnouncementResponse(this *ent.Announcement) *AdminAnnouncementResponse {
	return &AdminAnnouncementResponse{
		ID:         this.ID,
		Type:       this.Type,
		Message:    this.Message,
		CreatedAt:  this.CreatedAt.Format(time.RFC3339),
		ValidUntil: this.ValidUntil.Format(time.RFC3339),
	}
}

func responseAdminAnnouncementResponses(this []*ent.Announcement) []*AdminAnnouncementResponse {
	data := make([]*AdminAnnouncementResponse, len(this))
	for i, e := range this {
		data[i] = responseAdminAnnouncementResponse(e)
	}
	return data
}

// getAdminAnnouncement godoc
//
//	@Summary		List announcements
//	@Description	Get all announcements. Requires permission admin.announcement.view or admin.announcement.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		AdminAnnouncementResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/announcement [get]
func getAdminAnnouncement(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_ANNOUNCEMENT_VIEW, utils.PERMISSION_ANNOUNCEMENT_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		allEntires, err := db.Announcement.Query().All(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminAnnouncementResponses(allEntires))
	}
}

type CreateAnnouncement struct {
	Type       string `json:"type" enums:"ban,warning" validate:"required,announcement_type_validation"`
	Message    string `json:"message" example:"This is an example alert" validate:"required"`
	ValidUntil string `json:"validUntil" example:"2006-01-02T15:04:05Z07:00" validate:"omitempty,rfc3339"`
}

// postAdminAnnouncement godoc
//
//	@Summary		Create a new announcement
//	@Description	Requires permission admin.announcement.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		CreateAnnouncement	true	"-"
//	@Success		200		{object}	AdminAnnouncementResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/announcement [post]
func postAdminAnnouncement(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_ANNOUNCEMENT_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		req := new(CreateAnnouncement)
		if err := c.BodyParser(req); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if valid, err := handler.ValidateStruct(c, req); !valid {
			return err
		}

		var validUntil *time.Time
		if req.ValidUntil != "" {
			validUntilParsed, err := time.Parse(time.RFC3339, req.ValidUntil)
			if err != nil {
				validUntil = &validUntilParsed
			}
		}

		newEntry, err := db.Announcement.Create().SetType(req.Type).SetMessage(req.Message).SetNillableValidUntil(validUntil).Save(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminAnnouncementResponse(newEntry))
	}
}

// getAdminAnnouncementID godoc
//
//	@Summary		Get specific announcement
//	@Description	Get announcement by ID. Requires permission admin.announcement.view or admin.announcement.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Announcement ID"
//	@Success		200	{object}	AdminAnnouncementResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/announcement/{id} [get]
func getAdminAnnouncementID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_ANNOUNCEMENT_VIEW, utils.PERMISSION_ANNOUNCEMENT_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		foundEntry, err := db.Announcement.Query().Where(announcement.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminAnnouncementResponse(foundEntry))
	}
}

type UpdateAnnouncement struct {
	Type       string `json:"type" enums:"ban,warning" validate:"required,announcement_type_validation"`
	Message    string `json:"message" example:"This is an example alert" validate:"required"`
	ValidUntil string `json:"validUntil" example:"2006-01-02T15:04:05Z07:00" validate:"omitempty,rfc3339"`
}

// putAdminAnnouncementID godoc
//
//	@Summary		Update announcement
//	@Description	Update announcement by ID. Requires Permission admin.announcement.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		UpdateAnnouncement	true	"-"
//	@Param			id		path		int					true	"Announcement ID"
//	@Success		200		{object}	AdminAnnouncementResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/announcement/{id} [put]
func putAdminAnnouncementID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_ANNOUNCEMENT_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		req := new(UpdateAnnouncement)
		if err := c.BodyParser(req); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if valid, err := handler.ValidateStruct(c, req); !valid {
			return err
		}

		found, err := db.Announcement.Query().Where(announcement.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		var validUntil *time.Time
		if req.ValidUntil != "" {
			validUntilParsed, err := time.Parse(time.RFC3339, req.ValidUntil)
			if err != nil {
				validUntil = &validUntilParsed
			}
		}

		newFound, err := found.Update().SetType(req.Type).SetMessage(req.Message).SetNillableValidUntil(validUntil).Save(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminAnnouncementResponse(newFound))
	}
}

// deleteAdminAnnouncementID godoc
//
//	@Summary		Delete announcement
//	@Description	Delete announcement by ID. Requires permission admin.announcement.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Announcement ID"
//	@Success		200	{object}	handler.APIResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/announcement/{id} [delete]
func deleteAdminAnnouncementID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userId)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PERMISSION_ANNOUNCEMENT_EDIT}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		err = db.Announcement.DeleteOneID(pathID).Exec(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return handler.HandleSuccess(c)
	}
}
