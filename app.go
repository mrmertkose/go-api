package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/mrmertkose/go-api/app/database"
	"github.com/mrmertkose/go-api/routes"
	"log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file")
	}
}

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		log.Fatal("Not listen to port", err)
	}
}
