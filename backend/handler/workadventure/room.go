package workadventure

import (
	"context"
	"fmt"
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/announcement"
	"github.com/bo-mathadventure/admin/ent/maps"
	"github.com/bo-mathadventure/admin/ent/user"
	"github.com/bo-mathadventure/admin/handler"
	"github.com/bo-mathadventure/admin/utils"
	email "github.com/cameronnewman/go-emailvalidation/v3"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

func NewRoomHandler(app fiber.Router, ctx context.Context, db *ent.Client) {
	app.Get("/sameWorld", getSameWorld(ctx, db))
	app.Get("/access", getAccess(ctx, db))
}

type RoomSameWorld struct {
	RoomURL string `query:"roomUrl"`
}

func getSameWorld(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		qData := new(RoomSameWorld)
		if err := c.QueryParser(qData); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if qData.RoomURL == "" {
			return handler.HandleInsufficientData(c)
		}

		foundMaps, err := db.Maps.Query().All(ctx)
		if err != nil {
			return handler.HandleInternalError(c, err)
		}

		var res []string
		for _, m := range foundMaps {
			res = append(res, fmt.Sprintf("%s://%s%s", config.GetConfig().WorkadventureURLProtocol, config.GetConfig().WorkadventureURL, m.MapUrl))
		}

		return c.JSON(res)
	}
}

type RoomAccess struct {
	UserIdentifier  string   `query:"userIdentifier"`
	PlayURI         string   `query:"playUri"`
	IPAddress       string   `query:"ipAddress"`
	CharacterLayers []string `query:"characterLayers"`
}

func getAccess(ctx context.Context, db *ent.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		qData := new(RoomAccess)
		if err := c.QueryParser(qData); err != nil {
			return handler.HandleBodyParseError(c, err)
		}

		if qData.UserIdentifier == "" || qData.PlayURI == "" || qData.IPAddress == "" {
			return handler.HandleInsufficientData(c)
		}

		foundUsers, err := db.User.Query().Where(user.Or(user.UUIDEQ(qData.UserIdentifier), user.EmailEQ(email.Normalize(qData.UserIdentifier)))).All(ctx)
		if err != nil {
			return handler.HandleInvalidLogin(c)
		}

		playingMap := "/" + strings.Join(strings.Split(qData.PlayURI, "/")[3:], "/")
		mapURL := fmt.Sprintf("%s://%s%s", config.GetConfig().WorkadventureURLProtocol, config.GetConfig().WorkadventureURL, playingMap)

		foundMap, err := db.Maps.Query().Where(maps.MapUrlEQ(playingMap)).First(ctx)
		if err != nil {
			return handler.HandleInvalidLogin(c)
		}

		allAnouncements, err := db.Announcement.Query().Where(announcement.ValidUntilLTE(time.Now())).All(ctx)
		if err != nil {
			return handler.HandleInvalidLogin(c)
		}

		anouncementsList := []map[string]string{}
		for _, anouncement := range allAnouncements {
			anouncementsList = append(anouncementsList, map[string]string{
				"type":    anouncement.Type,
				"message": anouncement.Message,
			})
		}

		var foundUser *ent.User
		if len(foundUsers) > 0 {
			foundUser = foundUsers[0]
		}

		allTextures, err := db.Textures.Query().All(ctx)
		if err != nil {
			return handler.HandleInvalidLogin(c)
		}
		availableTexturesList := []map[string]string{}
		// make sure to preserve the texture order (given on characterLayers)
		for _, layer := range qData.CharacterLayers {
			for _, texture := range allTextures {
				if layer != texture.Texture {
					continue
				}
				if len(texture.Tags) > 0 && foundUser == nil {
					continue
				}
				if foundUser != nil && len(texture.Tags) > 0 && len(utils.ArrayIntersect(texture.Tags, foundUser.Tags)) == 0 {
					break
				}
				availableTexturesList = append(availableTexturesList, map[string]string{
					"id":    texture.Texture,
					"url":   strings.ReplaceAll(texture.URL, "%FRONTEND_URL%", config.GetConfig().FrontendURL),
					"layer": texture.Layer,
				})
				break
			}
		}

		resultData := map[string]interface{}{
			"userUuid":            qData.UserIdentifier,
			"email":               nil,
			"tags":                []string{},
			"textures":            availableTexturesList,
			"messages":            anouncementsList,
			"anonymous":           foundMap.PolicyNumber == 1,
			"visitCardUrl":        nil,
			"canEdit":             utils.CheckPermission(foundUser, utils.PERMISSION_MAP_EDITOR),
			"activatedInviteUser": false,
			"mucRooms": []interface{}{
				map[string]string{
					"name": foundMap.RoomName,
					"url":  mapURL,
					"type": "live",
				},
			},
		}

		if foundUser != nil {
			resultData["userUuid"] = foundUser.UUID
			resultData["email"] = foundUser.Email
			resultData["tags"] = foundUser.Tags
			resultData["username"] = foundUser.Username

			if foundUser.VCardURL != "" {
				resultData["visitCardUrl"] = foundUser.VCardURL
			}
		}

		return c.JSON(resultData)
	}
}
