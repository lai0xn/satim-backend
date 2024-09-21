package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Email  string
	Pass   string
	Dburl  string
	Dbpass string
)

func LoadENV() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Email = os.Getenv("EMAIL")
	Pass = os.Getenv("EMAIL_PASSWORD")
	Dburl = os.Getenv("DB_URL")
	Dbpass = os.Getenv("DB_PASS")


	log.Println("Config loaded successfully")
}
