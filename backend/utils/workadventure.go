package utils

import (
	"github.com/bo-mathadventure/admin/ent"
)

func UserCanAccessMap(userUUID string, foundUser *ent.User, foundMap *ent.Maps) bool {
	// null is acceptable here, because this is not meant for security, but for UX
	// WA uses call to /map Endpoint also without userId -> need to support that
	if userUUID == "" {
		return true
	}

	if foundMap.PolicyNumber != 1 {
		if foundMap.PolicyNumber == 2 {
			return foundUser != nil
		} else if foundMap.PolicyNumber == 3 {
			if len(foundMap.Tags) == 0 {
				return true
			}
			if foundUser == nil {
				return false
			}
			return len(ArrayIntersect(foundMap.Tags, foundUser.Tags)) >= 1
		}
	}
	return true
}
