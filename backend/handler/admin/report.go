package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/gofiber/fiber/v2"
	"time"
)

func NewAdminReportHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/", getAdminReport(ctx, db))
	app.Get("/:id", getAdminReportID(ctx, db))
	app.Put("/:id", putAdminReportID(ctx, db))
	app.Delete("/:id", deleteAdminReportID(ctx, db))
}

type AdminReportResponse struct {
	ID                  int       `json:"id"`
	ReportedUserComment string    `json:"reportedUserComment"`
	RoomUrl             string    `json:"roomUrl"`
	Hide                bool      `json:"hide"`
	CreatedAt           time.Time `json:"createdAt" example:"2006-01-02T15:04:05Z07:00"`
}

// getAdminReport godoc
//
//	@Summary		List reports
//	@Description	Get all reports. Requires permission admin.report.view or admin.report.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		AdminReportResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/report [get]
func getAdminReport(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_REPORT_VIEW || utils.PERMISSION_REPORT_EDIT
		return handler.HandleSuccess(c)
	}
}

// getAdminReportID godoc
//
//	@Summary		Get specific report
//	@Description	Get report by ID. Requires permission admin.report.view pr admin.report.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Report ID"
//	@Success		200	{object}	AdminReportResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/report/{id} [get]
func getAdminReportID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_REPORT_VIEW || utils.PERMISSION_REPORT_EDIT
		return handler.HandleSuccess(c)
	}
}

type UpdateReport struct {
}

// putAdminReportID godoc
//
//	@Summary		Update report
//	@Description	Update report by ID. Requires permission admin.report.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		UpdateReport	true	"-"
//	@Param			id		path		int				true	"Report ID"
//	@Success		200		{object}	AdminReportResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/report/{id} [put]
func putAdminReportID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_REPORT_EDIT
		return handler.HandleSuccess(c)
	}
}

// deleteAdminReportID godoc
//
//	@Summary		Delete report
//	@Description	Delete report by ID. Sets hide to true to keep history. Requires permission admin.report.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Report ID"
//	@Success		200	{object}	handler.APIResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/report/{id} [delete]
func deleteAdminReportID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// utils.PERMISSION_REPORT_EDIT
		return handler.HandleSuccess(c)
	}
}
