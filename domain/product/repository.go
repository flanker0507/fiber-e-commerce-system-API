package product

import (
	"fiber-e-commerce-system-API/domain/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	GetById(userID uint) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(id uint) (*models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) productRepository {
	return productRepository{db: db}
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetById(id uint) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *productRepository) Create(product models.Product) (models.Product, error) {
	err := r.db.Debug().Create(&product).Error
	return product, err
}

func (r *productRepository) UpdateProduct(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *productRepository) DeleteProduct(id uint) (*models.Product, error) {
	err := r.db.Unscoped().Delete(models.Product{}, id).Error
	if err != nil {
		return &models.Product{}, err
	}
	return &models.Product{}, err
}
