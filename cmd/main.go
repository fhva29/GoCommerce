package main

import (
	"log"

	"github.com/fhva29/GoCommerce/api"
	"github.com/fhva29/GoCommerce/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database
	config.ConnectDB()
	db := config.GetDB()

	// Gin
	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Register all routes
	api.RegisterRoutes(r, db)

	port := config.GetEnv("APP_PORT")
	r.Run(":" + port)
}
