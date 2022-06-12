package main

import (
	"github.com/joho/godotenv"
	"github.com/mrmertkose/go-api/database"
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
}
