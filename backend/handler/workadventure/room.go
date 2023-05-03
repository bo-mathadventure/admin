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
			return handler.HandleInsufficentData(c)
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
			return handler.HandleInsufficentData(c)
		}

		foundUser, err := db.User.Query().Where(user.UUIDEQ(qData.UserIdentifier)).First(ctx)
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

		var anouncementsList []map[string]string
		for _, anouncement := range allAnouncements {
			anouncementsList = append(anouncementsList, map[string]string{
				"type":    anouncement.Type,
				"message": anouncement.Message,
			})
		}

		allTextures, err := db.Textures.Query().All(ctx)
		if err != nil {
			return handler.HandleInvalidLogin(c)
		}
		var availableTexturesList []map[string]string
		// make sure to preserve the texture order (given on characterLayers)
		for _, layer := range qData.CharacterLayers {
			for _, texture := range allTextures {
				if layer != texture.Texture {
					continue
				}
				if len(texture.Tags) > 0 && len(utils.ArrayIntersect(texture.Tags, foundUser.Tags)) == 0 {
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
			"userUuid":            foundUser.UUID,
			"email":               foundUser.Email,
			"tags":                foundUser.Tags,
			"username":            foundUser.Username,
			"textures":            availableTexturesList,
			"messages":            anouncementsList,
			"anonymous":           foundMap.PolicyNumber == 1,
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

		if foundUser.VCardURL != "" {
			resultData["visitCardUrl"] = foundUser.VCardURL
		} else {
			resultData["visitCardUrl"] = nil
		}

		return c.JSON(resultData)
	}
}
