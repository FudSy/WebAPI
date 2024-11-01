package config

import (
	"os"
	"github.com/FudSy/webapi/pkg/models"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Initialize() {
	// initializing .env file
	err := godotenv.Load()
  	if err != nil {
    	logrus.Fatal("Error loading .env file")
  	}
	// initializing DB
	dsn := os.Getenv("DB")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Error opening postgres DB")
	}
	// migrating DB
	if err = DB.AutoMigrate(&models.Users{}); err != nil {
		logrus.Fatal("Error migrating DB")
	}
	
}