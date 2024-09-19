package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	JWTSecret string
)

func Load () {
	log.Println("Loading config")

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load config
	JWTSecret = os.Getenv("JWT_SECRET")
}