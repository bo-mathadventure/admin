package workadventure

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/handler"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
)

// NewReportHandler initialize routes for the given router
func NewReportHandler(ctx context.Context, app fiber.Router, db *ent.Client) {
	app.Post("/", getReport(ctx, db))
}

type reportQuery struct {
	ReportedUserUUID    string `json:"reportedUserUuid"`
	ReportedUserComment string `json:"reportedUserComment"`
	ReporterUserUIID    string `json:"reporterUserUuid"`
	RommURL             string `json:"reportWorldSlug"`
}

func getReport(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		qData := new(reportQuery)
		if err := c.BodyParser(qData); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if qData.ReportedUserUUID == "" || qData.ReporterUserUIID == "" || qData.RommURL == "" {
			return handler.HandleInsufficientData(c)
		}

		reportedUser, _ := db.User.Query().Where(user.Or(user.UUIDEQ(qData.ReportedUserUUID), user.EmailEQ(email.Normalize(qData.ReportedUserUUID)))).First(ctx)
		reporterUser, _ := db.User.Query().Where(user.Or(user.UUIDEQ(qData.ReporterUserUIID), user.EmailEQ(email.Normalize(qData.ReporterUserUIID)))).First(ctx)

		var reportedUserID *int
		if reportedUser != nil {
			reportedUserID = &reportedUser.ID
		}

		var reporterUserID *int
		if reporterUser != nil {
			reporterUserID = &reporterUser.ID
		}

		db.Report.Create().SetNillableReportedUserID(reportedUserID).SetNillableReporterUserID(reporterUserID).SetRoomUrl(qData.RommURL).SetReportedUserComment(qData.ReportedUserComment).Save(ctx)

		return handler.HandleSuccess(c)
	}
}
