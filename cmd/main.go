package main

import (
	"log"

	"github.com/fhva29/GoCommerce/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()

	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "ok",
		})
	})

	port := config.GetEnv("APP_PORT")
	r.Run(":" + port)
}
