package service

import (
	"challenge-2/models"
	"challenge-2/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	var productRepository = repository.ProductRepositoryMock{Mock: mock.Mock{}}
	var productServices = productService{repo: &productRepository}

	productRepository.Mock.On("GetProductById", uint(1)).Return(nil)

	product, err := productServices.GetProductById(1)

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "Error response has to be 'product not found'")
}

func TestProductServiceGetOneProductFound(t *testing.T) {
	var productRepository = repository.ProductRepositoryMock{Mock: mock.Mock{}}
	var productServices = productService{repo: &productRepository}

	products := models.Product{
		ID:          1,
		Title:       "Laptop",
		Description: "Lenovo A134",
	}

	productRepository.Mock.On("GetProductById", uint(1)).Return(products)

	product, err := productServices.GetProductById(1)

	assert.Nil(t, err)
	assert.NotNil(t, product)

	assert.Equal(t, &products, product, "result has to be a product with ID 1")
}

func TestProductServiceGetAllProductNotFound(t *testing.T) {
	var productRepository = repository.ProductRepositoryMock{
		Mock:     mock.Mock{},
		Products: []models.Product{},
		Error:    errors.New("product not found"),
	}
	var productServices = productService{repo: &productRepository}

	product, err := productServices.GetAllProduct()

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "products not found", err.Error(), "Error response has to be 'products not found'")
}

func TestProductServiceGetAllProductFound(t *testing.T) {
	var productRepository = repository.ProductRepositoryMock{
		Mock: mock.Mock{},
		Products: []models.Product{
			{
				ID:          uint(1),
				Title:       "Laptop",
				Description: "Lenovo A134",
			},
			{
				ID:          uint(2),
				Title:       "Macbook",
				Description: "Lenovo A134",
			}},
	}
	var productServices = productService{repo: &productRepository}

	productRepository.Mock.On("GetAllProduct").Return(productRepository.Products)

	result, err := productServices.GetAllProduct()

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, len(*result), len(productRepository.Products), "Expected %d products, but got %d", len(productRepository.Products), len(*result))
}
