package products

type ProductInput struct {
	Name     string  `json:"name" validate:"required"`
	Quantity int     `json:"quantity" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}

type ProductUpdate struct {
	ID       int
	Name     string  `json:"name" validate:"required"`
	Quantity int     `json:"quantity" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
}
