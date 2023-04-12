package service

import (
	"challenge-2/models"
	"challenge-2/repository"
)

func CreateProduct(product *models.Product) error {
	return repository.CreateProduct(product)
}

func UpdateProduct(product *models.Product) error {
	return repository.UpdateProduct(product)
}

func GetAllProducts() ([]models.Product, error) {
	return repository.GetAllProducts()
}

func GetProductById(productId uint) (models.Product, error) {
	return repository.GetProductById(productId)
}

func DeleteProductById(productId uint) error {
	return repository.DeleteProductById(productId)
}
