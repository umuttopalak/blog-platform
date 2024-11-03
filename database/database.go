package database

import (
	"blog-platform/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error

	dbType := os.Getenv("DB_TYPE")
	switch dbType {
	case "postgres":
		dsn := os.Getenv("DATABASE_URL")
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	}

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Reaction{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database connection successfully established")
}
