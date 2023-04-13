package service

import (
	"challenge-2/models"
	"challenge-2/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productServices = productService{repo: &productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("GetProductById", uint(1)).Return(nil)

	product, err := productServices.GetProductById(1)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "Error response has to be 'product not found'")
}

func TestProductServiceGetOneProductFound(t *testing.T) {
	product := models.Product{
		ID:          1,
		Title:       "Laptop",
		Description: "Lenovo A134",
		CreatedAt:   &time.Time{},
		UpdatedAt:   &time.Time{},
		UserID:      2,
		User:        nil,
	}

	productRepository.Mock.On("GetProductById", uint(1)).Return(product)

	result, err := productServices.GetProductById(1)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, product, result, "result has to be a product with ID 1")
}

func TestProductServiceGetAllProductNotFound(t *testing.T) {
	productRepository.Mock.On("GetAllProduct").Return(nil)

	product, err := productServices.GetAllProduct()

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "products not found", err.Error(), "Error response has to be 'products not found'")
}

func TestProductServiceGetAllProductFound(t *testing.T) {
	products := []models.Product{
		{
			ID:          uint(1),
			Title:       "Laptop",
			Description: "Lenovo A134",
		},
		{
			ID:          uint(2),
			Title:       "Macbook",
			Description: "Lenovo A134",
		},
	}

	productRepository.Mock.On("GetAllProduct").Return(products)

	result, err := productServices.GetAllProduct()

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, products, result, "result has to be a product with ID 1")
}
