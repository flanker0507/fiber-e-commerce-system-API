package user

import (
	"fiber-e-commerce-system-API/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	GetAll() ([]models.User, error)
	FindByEmail(email string) (models.User, error)
	GetById(userID uint) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRespository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var user []models.User

	err := r.db.Preload("Cart").Preload("Payments").Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, nil
	}

	return user, nil
}

func (r *userRepository) GetById(userID uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(user *models.User) (*models.User, error) {
	if err := r.db.Unscoped().Preload("Cart").Preload("Payments").Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
