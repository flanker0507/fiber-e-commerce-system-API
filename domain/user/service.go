package user

import (
	"errors"
	"fiber-e-commerce-system-API/domain/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(input UserInput) (models.User, error) {
	user := models.User{
		Name:    input.Name,
		Email:   input.Email,
		Role:    input.Role,
		Street:  input.Street,
		City:    input.City,
		State:   input.State,
		ZipCode: input.ZipCode,
		Country: input.Country,
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	return s.userRepository.Create(user)
}

func (s *UserService) GetAll() ([]models.User, error) {
	users, err := s.userRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) Login(input LoginInput) (models.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return user, nil
	}
	if user.ID == 0 {
		return user, errors.New("No User Found on That Email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *UserService) GetByID(userID uint) (*models.User, error) {
	return s.userRepository.GetById(userID)
}

func (s *UserService) UpdateUser(userID uint, input models.User) (*models.User, error) {
	user, err := s.userRepository.GetById(userID)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Role = input.Role
	user.Street = input.Street
	user.City = input.City
	user.State = input.State
	user.ZipCode = input.ZipCode
	user.Country = input.Country

	if input.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(passwordHash)
	}
	return s.userRepository.UpdateUser(user)
}

func (s *UserService) DeleteUser(userID uint) (*models.User, error) {
	user := &models.User{Model: gorm.Model{ID: userID}}
	deletedUser, err := s.userRepository.DeleteUser(user)
	if err != nil {
		return nil, err
	}
	return deletedUser, nil
}
