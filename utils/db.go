package utils

import (
	"app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	DB.AutoMigrate(&models.User{}, &models.Product{})
}

func CloseDB() {
	DB.Close()
}
