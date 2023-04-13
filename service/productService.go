package service

import (
	"challenge-2/models"
	"challenge-2/repository"
)

type ProductService interface {
	CreateProduct(photo *models.Product) error
	UpdateProduct(photo *models.Product, userId uint, id uint) error
	DeleteProduct(photo *models.Product, userId uint, id uint) error
	GetProductById(id uint) (*models.Product, error)
	GetAllProduct() (*[]models.Product, error)
}

type photoService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &photoService{repo: repo}
}

func (s *photoService) CreateProduct(photo *models.Product) error {
	return s.repo.CreateProduct(photo)
}

func (s *photoService) UpdateProduct(photo *models.Product, userId uint, id uint) error {
	photo.UserID = userId
	photo.ID = id
	return s.repo.UpdateProduct(photo)
}

func (s *photoService) DeleteProduct(photo *models.Product, userId uint, id uint) error {
	photo.UserID = userId
	photo.ID = id
	return s.repo.DeleteProduct(photo)
}

func (s *photoService) GetProductById(id uint) (*models.Product, error) {
	return s.repo.GetProductById(id)
}

func (s *photoService) GetAllProduct() (*[]models.Product, error) {
	return s.repo.GetAllProduct()
}
