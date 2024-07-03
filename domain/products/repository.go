package products

import (
	"fiber-e-commerce-system-API/domain/models"
	"gorm.io/gorm"
)

type Repository interface {
	Save(product models.Product) (models.Product, error)
	FindByID(ID int) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	FindAll() ([]models.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(product models.Product) (models.Product, error) {
	err := r.db.Debug().Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindByID(ID int) (models.Product, error) {
	var product models.Product

	err := r.db.Debug().Where("id = ?").Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Update(product models.Product) (models.Product, error) {
	err := r.db.Debug().Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindAll() ([]models.Product, error) {
	var product []models.Product

	err := r.db.Debug().Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
