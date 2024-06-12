package product

type ProductInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Qty         float64 `json:"qty" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}
