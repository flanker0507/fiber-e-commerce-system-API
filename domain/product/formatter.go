package product

import "fiber-e-commerce-system-API/domain/models"

type ProductRespon struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Qty         float64 `json:"qty"`
}

func FormatProduct(product models.Product) ProductRespon {
	return ProductRespon{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Qty:         product.Qty,
	}
}
