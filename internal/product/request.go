package product

import (
	"fmt"
	"strings"
)

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ValidationError struct {
	Errors map[string]string
}

func (v *ValidationError) Error() string {
	return "validation failed"
}

func (product *CreateProductRequest) Validate() map[string]string {
	errors := make(map[string]string)

	if err := validateStringField(product.Name, "name", 5, 25); err != "" {
		errors["name"] = err
	}

	if err := validateStringField(product.Description, "description", 10, 50); err != "" {
		errors["description"] = err
	}

	if product.Price < 0 {
		errors["price"] = "price must be greather than 0"
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func validateStringField(fieldValue string, fieldName string, min int, max int) string {
	value := strings.TrimSpace(fieldValue)

	if value == "" {
		return fmt.Sprintf("%s is required", fieldName)
	}

	if len(value) < min {
		return fmt.Sprintf("%s must be at least %d characters", fieldName, min)
	}

	if len(value) > max {
		return fmt.Sprintf("%s must be less than %d characters", fieldName, max)
	}

	return ""
}
