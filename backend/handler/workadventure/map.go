package workadventure

import (
	"context"
	"fmt"
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/utils"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
	"regexp"
	"time"
)

// NewMapHandler initialize routes for the given router
func NewMapHandler(ctx context.Context, app fiber.Router, db *ent.Client) {
	app.Get("/", getMap(ctx, db))
}

type mapQuery struct {
	PlayURI string `query:"playUri"`
	UserID  string `query:"userId"`
}

func getMap(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		qData := new(mapQuery)
		if err := c.QueryParser(qData); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		foundMap, mapURL, err := utils.GetMapFromPlayURL(ctx, db, qData.PlayURI)
		if qData.PlayURI == "" || mapURL == "" {
			return c.JSON(map[string]interface{}{
				"redirectUrl": config.GetConfig().WorkadventureStartRoomURL,
			})
		}
		// should not be triggered because mapURL is empty string
		if err != nil {
			return handler.HandleInvalidLogin(c)
		}

		if qData.UserID != "" {
			foundUser, _ := db.User.Query().Where(user.Or(user.UUIDEQ(qData.UserID), user.EmailEQ(email.Normalize(qData.UserID)))).First(ctx)
			if !utils.UserCanAccessMap(qData.UserID, foundUser, foundMap) {
				return handler.HandleInvalidLogin(c)
			}
		}

		result := map[string]interface{}{
			// "mapUrl":                  nil,
			// "wamUrl":                  nil,
			"authenticationMandatory": foundMap.PolicyNumber != 0,
			"group":                   nil,
			"mucRooms": []interface{}{
				map[string]string{
					"name": foundMap.RoomName,
					"url":  mapURL,
					"type": "live",
				},
			},
			// "contactPage":                nil,
			// "iframeAuthentication":       nil,
			// "opidLogoutRedirectUrl":      nil,
			"opidWokaNamePolicy": "force_opid",
			// "expireOn":  nil,
			"canReport": foundMap.CanReport,
			"editable":  true,
			// "loadingCowebsiteLogo":       nil,
			// "miniLogo":                   nil,
			// "loadingLogo":                nil,
			// "loginSceneLogo":             nil,
			// "showPoweredBy":              false,
			// "thirdParty":                 nil,
			// "metadata":                   nil,
			"roomName": foundMap.RoomName,
			// "pricingUrl":                 nil,
			"enableChat":                 foundMap.EnableChat,
			"enableChatUpload":           foundMap.EnableChatUpload,
			"enableChatOnlineList":       foundMap.EnableChatOnlineList,
			"enableChatDisconnectedList": foundMap.EnableChatDisconnectedList,
			// "metatags":                   nil,
			// "legals":                     nil,
			// "customizeWokaScene":         nil,
			// "backgroundColor":            nil,
			// "reportIssuesUrl":            nil,
			// "entityCollectionsUrls":      nil,
		}

		mapURLRegex := regexp.MustCompile(`\/_\/[^/]+\/(.+)`)
		wamURLRegex := regexp.MustCompile(`\/~\/(.+)`)
		if mapURLRegex.MatchString(foundMap.MapUrl) {
			matches := mapURLRegex.FindStringSubmatch(foundMap.MapUrl)
			result["mapUrl"] = fmt.Sprintf("%s://%s", config.GetConfig().WorkadventureURLProtocol, matches[1])
		} else if wamURLRegex.MatchString(foundMap.MapUrl) {
			matches := wamURLRegex.FindStringSubmatch(foundMap.MapUrl)
			// mapurl without .tmj extension!
			result["wamUrl"] = fmt.Sprintf("%s%s", config.GetConfig().MapStorageURL, matches[1])
			result["entityCollectionsUrls"] = []string{fmt.Sprintf("%s/entityCollections", config.GetConfig().MapStorageURL)}
		}

		if !foundMap.ExpireOn.IsZero() {
			result["expireOn"] = foundMap.ExpireOn.Format(time.RFC3339)
		}

		return c.JSON(result)
	}
}
