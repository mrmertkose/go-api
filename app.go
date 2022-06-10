package main

import (
	"github.com/joho/godotenv"
	"github.com/mrmertkose/go-api/database"
	"log"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

}
