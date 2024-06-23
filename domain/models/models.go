package models

import "time"

// User represents a user in the system
type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Role      string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Product represents an item with a name and quantity
type Product struct {
	ID       string
	Name     string
	Quantity int
	Price    float64
}

// CartItem represents an item in the cart
type CartItem struct {
	ProductID string
	Quantity  int
}

// Cart represents a shopping cart containing products
type Cart struct {
	UserID string
	Items  map[string]CartItem
}

// Transaction represents a completed purchase
type Transaction struct {
	ID     string
	UserID string
	Items  map[string]CartItem
	Total  float64
}
