package cart

import (
	"fiber-e-commerce-system-API/domain/models"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user models.User) (models.User, error)
	AddToCart(cartItem models.CartItem) (models.CartItem, error)
	GetCart(userID int) (models.Cart, error)
	UpdateProductQuantity(productID int, quantity int) error
	CreateCart(cart models.Cart) (models.Cart, error)
	ClearCart(userID int) error
	GetUserByID(userID int) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user models.User) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) AddToCart(cartItem models.CartItem) (models.CartItem, error) {
	err := r.db.Create(&cartItem).Error
	if err != nil {
		return cartItem, err
	}
	return cartItem, nil
}

func (r *repository) GetCart(userID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repository) UpdateProductQuantity(productID int, quantity int) error {
	var product models.Product
	err := r.db.Model(&product).Where("id = ?", productID).Update("quantity", gorm.Expr("quantity - ?", quantity)).Error
	return err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repository) ClearCart(userID int) error {
	var cart models.Cart
	err := r.db.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return err
	}

	err = r.db.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{}).Error
	return err
}

func (r *repository) GetUserByID(userID int) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", userID).First(&user).Error
	return user, err
}
