package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadAPIKey loads the Gemini API key from the .env file
func LoadAPIKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY is not set in .env")
	}
	return apiKey
}
