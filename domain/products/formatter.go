package products

import "fiber-e-commerce-system-API/domain/models"

type ProductFormatter struct {
	ID       int     `json:"id" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Quantity int     `json:"quantity" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

func FormatProduct(product models.Product) ProductFormatter {
	formatter := ProductFormatter{
		ID:       product.ID,
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}

	return formatter
}
