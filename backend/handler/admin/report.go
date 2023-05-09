package admin

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/report"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

// NewAdminReportHandler initialize routes for the given router
func NewAdminReportHandler(ctx context.Context, app fiber.Router, db *ent.Client) {
	app.Get("/", getAdminReport(ctx, db))
	app.Get("/:id", getAdminReportID(ctx, db))
	app.Put("/:id", putAdminReportID(ctx, db))
	app.Delete("/:id", deleteAdminReportID(ctx, db))
}

type adminReportResponse struct {
	ID                  int    `json:"id"`
	ReportedUserComment string `json:"reportedUserComment"`
	RoomURL             string `json:"roomUrl"`
	Hide                bool   `json:"hide"`
	CreatedAt           string `json:"createdAt" example:"2006-01-02T15:04:05Z07:00"`
}

func responseAdminReportResponse(this *ent.Report) *adminReportResponse {
	return &adminReportResponse{
		ID:                  this.ID,
		ReportedUserComment: this.ReportedUserComment,
		RoomURL:             this.RoomUrl,
		Hide:                this.Hide,
		CreatedAt:           this.CreatedAt.Format(time.RFC3339),
	}
}

func responseAdminReportResponses(this []*ent.Report) []*adminReportResponse {
	data := make([]*adminReportResponse, len(this))
	for i, e := range this {
		data[i] = responseAdminReportResponse(e)
	}
	return data
}

// getAdminReport godoc
//
//	@Summary		List reports
//	@Description	Get all reports. Requires permission admin.report.view or admin.report.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		adminReportResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/report [get]
func getAdminReport(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userID)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PermissionReportView, utils.PermissionReportEdit}) {
			return handler.HandleInvalidPermissions(c)
		}

		allEntires, err := db.Report.Query().All(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminReportResponses(allEntires))
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
//	@Success		200	{object}	adminReportResponse
//	@Failure		400	{object}	handler.APIResponse
//	@Failure		401	{object}	handler.APIResponse
//	@Failure		404	{object}	handler.APIResponse
//	@Failure		500	{object}	handler.APIResponse
//	@Router			/system/admin/report/{id} [get]
func getAdminReportID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userID)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PermissionReportView, utils.PermissionReportEdit}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		foundEntry, err := db.Report.Query().Where(report.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminReportResponse(foundEntry))
	}
}

type updateReport struct {
	// fixme add at least staff comments
}

// putAdminReportID godoc
//
//	@Summary		Update report
//	@Description	Update report by ID. Requires permission admin.report.edit
//	@Security		ApiKeyAuth
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			params	body		updateReport	true	"-"
//	@Param			id		path		int				true	"Report ID"
//	@Success		200		{object}	adminReportResponse
//	@Failure		400		{object}	handler.APIResponse
//	@Failure		401		{object}	handler.APIResponse
//	@Failure		404		{object}	handler.APIResponse
//	@Failure		500		{object}	handler.APIResponse
//	@Router			/system/admin/report/{id} [put]
func putAdminReportID(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userID)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PermissionReportEdit}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		req := new(updateReport)
		if err := c.BodyParser(req); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if valid, err := handler.ValidateStruct(c, req); !valid {
			return err
		}

		found, err := db.Report.Query().Where(report.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		newFound, err := found.Update().Save(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return c.JSON(responseAdminReportResponse(newFound))
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
		jwtUser := c.Locals("user").(*jwt.Token)
		claims := jwtUser.Claims.(jwt.MapClaims)
		userID := int(claims["id"].(float64))

		thisUser, err := db.User.Query().WithGroups().Where(user.ID(userID)).First(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		if !utils.CheckPermissionAny(thisUser, []string{utils.PermissionReportEdit}) {
			return handler.HandleInvalidPermissions(c)
		}

		pathID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return handler.HandleInvalidID(c)
		}

		found, err := db.Report.Query().Where(report.ID(pathID)).First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return handler.HandleNotFound(c)
			}
			return handler.HandleInternalError(c, err)
		}

		_, err = found.Update().SetHide(true).Save(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		return handler.HandleSuccess(c)
	}
}
