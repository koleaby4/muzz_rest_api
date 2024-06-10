package config

import (
	"github.com/lpernett/godotenv"
	"log"
	"os"
)

// GetConfig reads the .env file and returns the value of the key
func GetConfig(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Error getting %s from .env file", key)
	}
	return val
}
