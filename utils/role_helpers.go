package utils

import "blog-platform/models"

func HasRole(user models.User, roleName string) bool {
	for _, role := range user.Roles {
		if role.Name == roleName {
			return true
		}
	}
	return false
}
