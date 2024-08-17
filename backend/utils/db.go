package utils

import (
	"backend/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {
	var err error
	dsn := os.Getenv("DATABASE_LOCAL_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error loading DB.")
	}

	migrateError := DB.AutoMigrate(&models.User{})
	if migrateError != nil {
		panic("Error migration user")
	}

	fmt.Println("Connected to DB.")

	return DB
}
