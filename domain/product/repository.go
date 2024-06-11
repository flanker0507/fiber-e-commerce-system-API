package product

import (
	"fiber-e-commerce-system-API/domain/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	GetById(userID uint) (*models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	DeleteUser(user *models.Product) (*models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var product []models.Product

	err := r.db.Preload("Cart").Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil

}
