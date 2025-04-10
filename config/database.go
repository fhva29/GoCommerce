package config

import (
	"fmt"
	"log"

	internal "github.com/fhva29/GoCommerce/internal/product"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	db_url := GetEnv("DB_URL")
	connection, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database!")
	err = connection.AutoMigrate(&internal.Product{})
	if err != nil {
		log.Fatal(err)
	}

	db = connection
}

func GetDB() *gorm.DB {
	return db
}
