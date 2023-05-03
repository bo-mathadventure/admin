package utils

import (
	"github.com/bo-mathadventure/admin/ent"
	"strings"
)

func CheckPermission(user *ent.User, requiredPermission string) bool {
	if user == nil {
		return false
	}
	for _, permission := range user.Permissions {
		if matchPermission(permission, requiredPermission) {
			return true
		}
	}
	return false
}

func ComparePermission(testPermission string, requiredPermission string) bool {
	if testPermission == "" || requiredPermission == "" {
		return false
	}
	return matchPermission(strings.ToLower(testPermission), strings.ToLower(requiredPermission))
}

func matchPermission(first string, second string) bool {
	// If we reach at the end of both strings, we are done
	if len(first) == 0 && len(second) == 0 {
		return true
	}

	// Make sure that the characters after '*'
	// are present in second string.
	// This function assumes that the first
	// string will not contain two consecutive '*'
	if len(first) > 1 && first[0] == '*' && len(second) == 0 {
		return false
	}

	// If the first string contains '?',
	// or current characters of both strings match
	if (len(first) > 1 && first[0] == '?') ||
		(len(first) != 0 && len(second) != 0 && first[0] == second[0]) {
		return matchPermission(first[1:], second[1:])
	}

	// If there is *, then there are two possibilities
	// a) We consider current character of second string
	// b) We ignore current character of second string.
	if len(first) > 0 && first[0] == '*' {
		return matchPermission(first[1:], second) || matchPermission(first, second[1:])
	}

	return false
}
