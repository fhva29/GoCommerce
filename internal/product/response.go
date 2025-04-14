package product

type ProductResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func ToProductResponse(product Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}

func ToProductResponseList(products []Product) []ProductResponse {
	var result []ProductResponse
	for _, product := range products {
		result = append(result, ToProductResponse(product))
	}

	return result
}
