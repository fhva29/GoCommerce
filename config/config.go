package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Variable %s not present on environment", key)
	}

	return value
}
