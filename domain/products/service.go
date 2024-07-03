package products

import (
	"errors"
	"fiber-e-commerce-system-API/domain/models"
)

type Service interface {
	CreateProduct(input ProductInput) (models.Product, error)
	GetProductByID(ID int) (models.Product, error)
	GetAllProduct() ([]models.Product, error)
	UpdateUser(input ProductUpdate) (models.Product, error)
}

type service struct {
	repo Repository
}

func NewServiceProduct(repo Repository) *service {
	return &service{repo}
}

func (s *service) CreateProduct(input ProductInput) (models.Product, error) {
	product := models.Product{}
	product.Name = input.Name
	product.Quantity = input.Quantity
	product.Price = input.Price

	newProduct, err := s.repo.Save(product)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (s *service) GetProductByID(ID int) (models.Product, error) {
	product, err := s.repo.FindByID(ID)
	if err != nil {
		return product, err
	}

	if product.ID == 0 {
		return product, errors.New("No product found on that ID")
	}

	return product, nil
}

func (s *service) GetAllProduct() ([]models.Product, error) {
	product, err := s.repo.FindAll()
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *service) UpdateUser(input ProductUpdate) (models.Product, error) {
	product, err := s.repo.FindByID(input.ID)
	if err != nil {
		return product, err
	}

	product.Name = input.Name
	product.Quantity = input.Quantity
	product.Price = input.Price

	update, err := s.repo.Update(product)
	if err != nil {
		return update, err
	}

	return update, nil
}
