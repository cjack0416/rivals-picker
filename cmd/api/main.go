package main

import (
	"context"
	"os"

	"github.com/cjack0416/rivals-picker/internal/handlers"
	"github.com/cjack0416/rivals-picker/internal/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	log.Info("Connecting to database")

	dbURI := os.Getenv("DB_URI")
	conn, err := pgx.Connect(context.Background(), dbURI)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	log.Info("Successfully connected to database")

	app := fiber.New()

	api := app.Group("/api")

	heroPicker := api.Group("/hero-picker")

	tools.SetDatabaseConn(conn)

	heroPicker.Get("/competitive", handlers.CompetitivePicker)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}