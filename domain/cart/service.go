package cart

import (
	"errors"
	"fiber-e-commerce-system-API/domain/models"
)

type Service interface {
	AddItemToCart(userID int, productID int, quantity int) (models.CartItem, error)
	GetUserCart(userID int) (models.Cart, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) AddItemToCart(userID int, productID int, quantity int) (models.CartItem, error) {
	// Ensure the user exists
	_, err := s.repo.GetUserByID(userID)
	if err != nil {
		return models.CartItem{}, errors.New("user does not exist")
	}

	// Ensure the cart exists for the user
	cart, err := s.repo.GetCart(userID)
	if err != nil && err.Error() == "record not found" {
		cart = models.Cart{
			UserID: userID,
		}
		cart, err = s.repo.CreateCart(cart)
		if err != nil {
			return models.CartItem{}, err
		}
	} else if err != nil {
		return models.CartItem{}, err
	}

	// Decrease product quantity
	err = s.repo.UpdateProductQuantity(productID, quantity)
	if err != nil {
		return models.CartItem{}, err
	}

	// Add item to cart
	cartItem := models.CartItem{
		CartID:    cart.ID,
		ProductID: productID,
		Quantity:  quantity,
	}

	newCartItem, err := s.repo.AddToCart(cartItem)
	if err != nil {
		return newCartItem, err
	}
	return newCartItem, nil
}

func (s *service) GetUserCart(userID int) (models.Cart, error) {
	return s.repo.GetCart(userID)
}
