package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Email string
	Pass string
)

func Load () {
	log.Println("Loading config")

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load config
	Email = os.Getenv("EMAIL")
	Pass = os.Getenv("EMAIL_PASSWORD")
}
