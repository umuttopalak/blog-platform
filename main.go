package main

import (
	"blog-platform/database"
	"blog-platform/routes"
	"blog-platform/utils"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	database.InitDatabase()

	utils.SeedRoles(database.DB)

	router := routes.SetupRouter()

	err = router.Run(":8080")
	if err != nil {
		log.Println("Err while running")
	}
}
