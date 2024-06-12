package product

import (
	"fiber-e-commerce-system-API/domain/models"
)

type ProductService interface {
	GetALlProduct() ([]ProductRespon, error)
	GetProductByID(id uint) (ProductRespon, error)
	CreateProduct(input ProductInput) (ProductRespon, error)
	UpdateProduct(id uint, input ProductInput) (ProductRespon, error)
	DeleteProduct(id uint) (*models.Product, error)
}

type productService struct {
	repo productRepository
}

func NewProductService(repo productRepository) productService {
	return productService{repo: repo}
}

func (s *productService) GetAllProduct() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) GetProductByID(id uint) (ProductRespon, error) {
	input, err := s.repo.GetById(id)
	if err != nil {
		return ProductRespon{}, err
	}

	return ProductRespon{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Qty:         input.Qty,
		Price:       input.Price,
	}, nil
}

func (s *productService) CreateProduct(input ProductInput) (ProductRespon, error) {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Qty:         input.Qty,
		Price:       input.Price,
	}

	newProduct, err := s.repo.Create(product)
	if err != nil {
		return ProductRespon{}, err
	}

	return ProductRespon{
		ID:          newProduct.ID,
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Qty:         newProduct.Qty,
		Price:       newProduct.Price,
	}, err

}

func (s *productService) UpdateProduct(id uint, input ProductInput) (ProductRespon, error) {
	product, err := s.repo.GetById(id)
	if err != nil {
		return ProductRespon{}, err
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Qty = input.Qty

	updatedProduct, err := s.repo.UpdateProduct(product)
	if err != nil {
		return ProductRespon{}, err
	}

	return ProductRespon{
		ID:          updatedProduct.ID,
		Name:        updatedProduct.Name,
		Description: updatedProduct.Description,
		Qty:         updatedProduct.Qty,
		Price:       updatedProduct.Price,
	}, err
}

func (s *productService) DeleteProduct(id uint) (*models.Product, error) {
	return s.repo.DeleteProduct(id)
}
