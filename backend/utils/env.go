package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	os := godotenv.Load()
	if os != nil {
		log.Fatal("Error loading .env file", os)
	}
}
