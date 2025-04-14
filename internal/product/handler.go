package product

import (
	"net/http"

	"github.com/fhva29/GoCommerce/internal/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) GetAll(ctx *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		response.SendError(ctx, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.SendSuccess(ctx, ToProductResponseList(products))
}

func (h *Handler) Create(ctx *gin.Context) {
	var req CreateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.SendError(ctx, http.StatusBadRequest, "invalid request", nil)
		return
	}

	product, err := h.service.CreateProduct(req)
	if err != nil {
		if ve, ok := err.(*ValidationError); ok {
			response.SendError(ctx, http.StatusBadRequest, "validation error", ve.Errors)
			return
		}

		response.SendError(ctx, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.SendSuccess(ctx, ToProductResponse(product))
}
