package product

type ProductInput struct {
	Name       string  `json:"name" binding:"required"`
	Descrition string  `json:"descrition" binding:"required"`
	Price      float64 `json:"price" binding:"required"`
}
