package api

import (
	"github.com/fhva29/GoCommerce/internal/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	RegisterProductRoutes(r, db)
}

func RegisterProductRoutes(r *gin.Engine, db *gorm.DB) {
	repo := product.NewRepository(db)
	service := product.NewService(repo)
	handler := product.NewHandler(service)

	r.GET("/products", handler.GetAll)
}
