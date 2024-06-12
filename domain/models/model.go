package models

import (
	"gorm.io/gorm"
)

type CartProduct struct {
	CartID    uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
}

type Cart struct {
	gorm.Model
	UserID   uint
	Products []Product `gorm:"many2many:cart_products"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string `json:"-"`
	Role     string
	Street   string
	City     string
	State    string
	ZipCode  string
	Country  string
	Cart     Cart
	Payments []Payment `gorm:"foreignKey:UserID"`
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Qty         float64
	Carts       []Cart `gorm:"many2many:cart_product"`
}

type Payment struct {
	gorm.Model
	UserID        uint
	Amount        float64
	PaymentMethod string
	status        string
}
