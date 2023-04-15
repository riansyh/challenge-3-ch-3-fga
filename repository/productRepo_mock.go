package repository

import (
	"challenge-2/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock     mock.Mock
	Products []models.Product
	Error    error
}

func (repository *ProductRepositoryMock) GetProductById(id uint) *models.Product {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(models.Product)

	return &product
}

func (repository *ProductRepositoryMock) CreateProduct(product *models.Product) error {
	return nil
}

func (repository *ProductRepositoryMock) UpdateProduct(product *models.Product) error {
	return nil
}

func (repository *ProductRepositoryMock) DeleteProduct(product *models.Product) error {
	return nil
}

func (repository *ProductRepositoryMock) GetAllProduct() *[]models.Product {
	if repository.Error != nil {
		return nil
	}

	return &repository.Products
}
