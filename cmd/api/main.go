package main

import (
	"log"
	"os"

	"github.com/cjack0416/rivals-picker/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	app := fiber.New()

	api := app.Group("/api")

	heroPicker := api.Group("/hero-picker")

	heroPicker.Get("/competitive", handlers.CompetitivePicker)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}