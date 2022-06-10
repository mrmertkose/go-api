package database

import (
	"github.com/mrmertkose/go-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:rootroot@/databaseName"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database!", err)
	}

	connection.AutoMigrate(&models.User{})
}
