package utils

import (
	"github.com/bo-mathadventure/admin/ent"
	"sort"
)

func CheckPermissionAny(user *ent.User, requiredPermissions []string) bool {
	if user == nil {
		return false
	}
	return len(ArrayIntersect(CombinePermissions(user), requiredPermissions)) > 0
}

func CheckPermission(user *ent.User, requiredPermission string) bool {
	return CheckPermissionAny(user, []string{requiredPermission})
}

func CombinePermissions(user *ent.User) []string {
	if user.Edges.Groups == nil {
		return user.Permissions
	}

	var all []string
	for _, group := range user.Edges.Groups {
		all = append(all, group.Permissions...)
	}

	sort.Strings(all)
	return all
}

func CombineTags(user *ent.User) []string {
	if user.Edges.Groups == nil {
		return user.Tags
	}

	var all []string
	for _, group := range user.Edges.Groups {
		all = append(all, group.Tags...)
	}

	sort.Strings(all)
	return all
}
