// utils/seed.go
package utils

import (
	"blog-platform/models"
	"log"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{Name: "Admin", Description: "Has full access to all resources and settings"},
		{Name: "Editor", Description: "Can edit and publish content"},
		{Name: "Author", Description: "Can create and edit own content"},
		{Name: "Reader", Description: "Can view content only"},
	}

	for _, role := range roles {
		if err := db.Where("name = ?", role.Name).FirstOrCreate(&role).Error; err != nil {
			log.Printf("Role %s could not be created: %v", role.Name, err)
		}
	}
}
