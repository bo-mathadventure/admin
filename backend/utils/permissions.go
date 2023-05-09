package utils

import (
	"github.com/bo-mathadventure/admin/ent"
	"sort"
)

// CheckPermissionAny checks if any permission given in the array is available
func CheckPermissionAny(user *ent.User, requiredPermissions []string) bool {
	if user == nil {
		return false
	}
	return len(ArrayIntersect(CombinePermissions(user), requiredPermissions)) > 0
}

// CheckPermission same as CheckPermissionAny but only uses a single string
func CheckPermission(user *ent.User, requiredPermission string) bool {
	return CheckPermissionAny(user, []string{requiredPermission})
}

// CombinePermissions combines permissions of all groups from the user
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

// CombineTags same as CombinePermissions but for Tags
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
