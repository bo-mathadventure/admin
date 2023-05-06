package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
	"time"
)

func NewAdminAnnouncementHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminAnnouncement(ctx, db))
	app.Post("/", postAdminAnnouncement(ctx, db))
	app.Get("/:id", getAdminAnnouncementID(ctx, db))
	app.Put("/:id", putAdminAnnouncementID(ctx, db))
	app.Delete("/:id", deleteAdminAnnouncementID(ctx, db))
}

type AdminAnnouncementResponse struct {
	ID         int       `json:"id"`
	Type       string    `json:"type" enums:"ban,warning"`
	Message    string    `json:"message" example:"This is an example alert"`
	CreatedAt  time.Time `json:"createdAt" example:"2006-01-02T15:04:05Z07:00"`
	ValidUntil time.Time `json:"validUntil" example:"2006-01-02T15:04:05Z07:00" validate:"optional"`
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
		// utils.PERMISSION_ANNOUNCEMENT_VIEW || utils.PERMISSION_ANNOUNCEMENT_EDIT
		return handler.HandleSuccess(c)
	}
}

type CreateAnnouncement struct {
	Type       string    `json:"type" enums:"ban,warning"`
	Message    string    `json:"message" example:"This is an example alert"`
	ValidUntil time.Time `json:"validUntil" example:"2006-01-02T15:04:05Z07:00" validate:"optional"`
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
		// utils.PERMISSION_ANNOUNCEMENT_EDIT
		return handler.HandleSuccess(c)
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
		// utils.PERMISSION_ANNOUNCEMENT_VIEW || utils.PERMISSION_ANNOUNCEMENT_EDIT
		return handler.HandleSuccess(c)
	}
}

type UpdateAnnouncement struct {
	Type       string    `json:"type" enums:"ban,warning"`
	Message    string    `json:"message" example:"This is an example alert"`
	ValidUntil time.Time `json:"validUntil" example:"2006-01-02T15:04:05Z07:00" validate:"optional"`
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
		// utils.PERMISSION_ANNOUNCEMENT_EDIT
		return handler.HandleSuccess(c)
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
		// utils.PERMISSION_ANNOUNCEMENT_EDIT
		return handler.HandleSuccess(c)
	}
}
