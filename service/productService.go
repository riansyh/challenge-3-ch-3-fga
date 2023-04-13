package service

import (
	"challenge-2/models"
	"challenge-2/repository"
	"errors"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product, userId uint, id uint) error
	DeleteProduct(product *models.Product, userId uint, id uint) error
	GetProductById(id uint) (*models.Product, error)
	GetAllProduct() (*[]models.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *productService) UpdateProduct(product *models.Product, userId uint, id uint) error {
	product.UserID = userId
	product.ID = id
	return s.repo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(product *models.Product, userId uint, id uint) error {
	product.UserID = userId
	product.ID = id
	return s.repo.DeleteProduct(product)
}

func (s *productService) GetProductById(id uint) (*models.Product, error) {
	product := s.repo.GetProductById(id)

	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (s *productService) GetAllProduct() (*[]models.Product, error) {
	products := s.repo.GetAllProduct()

	if products == nil {
		return nil, errors.New("products not found")
	}

	return products, nil
}
