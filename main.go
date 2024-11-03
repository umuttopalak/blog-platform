package main

import (
	"blog-platform/database"
	"blog-platform/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	database.InitDatabase()

	router := routes.SetupRouter()

	err = router.Run(":8080")
	if err != nil {
		log.Println("Err while running")
	}
}
