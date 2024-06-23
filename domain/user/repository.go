package user

import (
	"fiber-e-commerce-system-API/domain/models"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByID(ID int) (models.User, error)
	Update(user models.User) (models.User, error)
	FindAll() ([]models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user models.User) (models.User, error) {
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Debug().Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(id int) (models.User, error) {
	var user models.User

	err := r.db.Debug().Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user models.User) (models.User, error) {
	err := r.db.Debug().Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAll() ([]models.User, error) {
	var user []models.User

	err := r.db.Debug().Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
