package cart

import (
	"fiber-e-commerce-system-API/domain/models"
)

type AddProductInput struct {
	//UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	User      models.User
}
