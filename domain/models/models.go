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
	ID       int
	Name     string
	Quantity int
	Price    float64
}

//// CartItem represents an item in the cart
//type CartItem struct {
//	ProductID int
//	Quantity  int
//}
//
//// Cart represents a shopping cart containing products
//type Cart struct {
//	UserID int
//	Items  map[string]CartItem
//}
//
//package models
//
//import "time"

// Cart represents a shopping cart containing products
type Cart struct {
	ID        int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []CartItem `gorm:"foreignKey:CartID"`
	User      User
}

//// CartItem represents an item in the cart
//type CartItem struct {
//	ID        int
//	CartID    int
//	ProductID int
//	Quantity  int
//	Product   Product `gorm:"foreignKey:ProductID"`
//}

type CartItem struct {
	ID        int
	CartID    int
	ProductID int
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Transaction struct {
	ID         int
	CartID     int
	UserID     int
	Total      float64
	Status     string
	PaymentURL string
	Items      []TransactionItem `gorm:"foreignKey:TransactionID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type TransactionItem struct {
	ID            int
	TransactionID int
	ProductID     int
	Quantity      int
	Product       Product `gorm:"foreignKey:ProductID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
