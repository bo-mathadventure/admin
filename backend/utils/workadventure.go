package utils

import (
	"context"
	"fmt"
	"github.com/bo-mathadventure/admin/config"
	"github.com/bo-mathadventure/admin/ent"
	"github.com/bo-mathadventure/admin/ent/maps"
	"strings"
	"time"
)

func UserCanAccessMap(userUUID string, foundUser *ent.User, foundMap *ent.Maps) bool {
	// null is acceptable here, because this is not meant for security, but for UX
	// WA uses call to /map Endpoint also without userId -> need to support that
	if userUUID == "" {
		return true
	}

	if foundMap.PolicyNumber == 0 {
		return true
	} else if foundMap.PolicyNumber == 1 {
		return foundUser != nil
	} else if foundMap.PolicyNumber == 2 {
		if len(foundMap.Tags) == 0 {
			return true
		}
		if foundUser == nil {
			return false
		}
		return len(ArrayIntersect(foundMap.Tags, foundUser.Tags)) >= 1
	}
	return true
}

func GetMapFromPlayURL(ctx context.Context, db *ent.Client, playURL string) (*ent.Maps, string, error) {
	playingMap := "/" + strings.Join(strings.Split(playURL, "/")[3:], "/")
	mapURL := fmt.Sprintf("%s://%s%s", config.GetConfig().WorkadventureURLProtocol, config.GetConfig().WorkadventureURL, playingMap)
	foundMap, err := db.Maps.Query().Where(maps.MapUrlEQ(playingMap)).Where(maps.Or(maps.ExpireOnIsNil(), maps.ExpireOnGTE(time.Now()))).First(ctx)
	if err != nil {
		return nil, "", err
	}
	return foundMap, mapURL, nil
}
