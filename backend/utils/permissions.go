package utils

import (
	"github.com/bo-mathadventure/admin/ent"
)

func CheckPermissionAny(user *ent.User, requiredPermissions []string) bool {
	if user == nil {
		return false
	}
	return len(ArrayIntersect(user.Permissions, requiredPermissions)) > 0
}

func CheckPermission(user *ent.User, requiredPermission string) bool {
	return CheckPermissionAny(user, []string{requiredPermission})
}
