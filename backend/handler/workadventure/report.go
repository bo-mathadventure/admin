package workadventure

import (
	"context"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/handler"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
)

func NewReportHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Post("/report", getReport(ctx, db))
}

type ReportQuery struct {
	ReportedUserUUID    string `query:"reportedUserUuid"`
	ReportedUserComment string `query:"reportedUserComment"`
	ReporterUserUIID    string `query:"reporterUserUuid"`
	RommURL             string `query:"roomUrl"`
}

func getReport(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		qData := new(ReportQuery)
		if err := c.QueryParser(qData); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if qData.ReportedUserUUID == "" || qData.ReporterUserUIID == "" || qData.RommURL == "" {
			return handler.HandleInsufficentData(c)
		}

		reportedUser, _ := db.User.Query().Where(user.Or(user.UUIDEQ(qData.ReportedUserUUID), user.EmailEQ(email.Normalize(qData.ReportedUserUUID)))).First(ctx)
		reporterUser, _ := db.User.Query().Where(user.Or(user.UUIDEQ(qData.ReporterUserUIID), user.EmailEQ(email.Normalize(qData.ReporterUserUIID)))).First(ctx)

		if reportedUser == nil || reporterUser == nil {
			return handler.HandleSuccess(c)
		}

		db.Report.Create().SetReportedUser(reportedUser).SetReporterUser(reporterUser).SetRoomUrl(qData.RommURL).SetReportedUserComment(qData.ReportedUserComment).Save(ctx)

		return handler.HandleSuccess(c)
	}
}
