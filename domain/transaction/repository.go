package transaction

import (
	"fiber-e-commerce-system-API/domain/models"
	"gorm.io/gorm"
)

type Repository interface {
	Save(transaction models.Transaction) (models.Transaction, error)
	GetByID(transactionID int) (models.Transaction, error)
	GetByUserID(userID int) ([]models.Transaction, error)
	Update(transaction models.Transaction) (models.Transaction, error)
	FindAll() ([]models.Transaction, error)
	GetCartItems(cartID int) ([]models.CartItem, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) GetByID(transactionID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Items.Product").Where("id = ?", transactionID).First(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) GetByUserID(userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("user_id = ?", userID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) Update(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) FindAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) GetCartItems(cartID int) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	err := r.db.Where("cart_id = ?", cartID).Preload("Product").Find(&cartItems).Error
	if err != nil {
		return cartItems, err
	}
	return cartItems, nil
}
