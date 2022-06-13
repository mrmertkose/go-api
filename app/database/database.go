package database

import (
	"fmt"
	"github.com/mrmertkose/go-api/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connect() {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	connection, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database!", err)
	}

	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration Error!", err)
	}
}
