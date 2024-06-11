package product

type ProductRespon struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Descrition string  `json:"descrition"`
	Price      float64 `json:"price"`
}
