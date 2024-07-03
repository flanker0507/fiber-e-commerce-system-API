package transaction

import "fiber-e-commerce-system-API/domain/models"

type CreateTransactionInput struct {
	CartID int `json:"cart_id"`
	UserID int `json:"user_id"`
	Total  float64
	User   models.User
}

type TransactionNotificationInput struct {
	OrderID           string
	PaymentType       string
	TransactionStatus string
	FraudStatus       string
}
