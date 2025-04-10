package config

import (
	"log"
	"os"
)

func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Variable %s not present on environment", key)
	}

	return value
}
