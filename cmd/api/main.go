package main

import (
	"context"
	"os"

	"github.com/cjack0416/rivals-picker/internal/handlers"
	"github.com/cjack0416/rivals-picker/internal/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
			os.Exit(1)
		}
	}

	dbURI := os.Getenv("DB_URI")
	conn, err := tools.InitializeDatabaseConn(dbURI)
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close(context.Background())

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