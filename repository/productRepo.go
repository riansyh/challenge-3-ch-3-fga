package repository

import (
	"challenge-2/database"
	"challenge-2/models"
)

func CreateProduct(product *models.Product) error {
	db := database.GetDB()
	return db.Debug().Create(product).Error
}

func UpdateProduct(product *models.Product) error {
	db := database.GetDB()
	return db.Model(product).Updates(models.Product{Title: product.Title, Description: product.Description}).Error
}

func GetAllProducts() ([]models.Product, error) {
	db := database.GetDB()
	products := []models.Product{}
	err := db.Find(&products).Error
	return products, err
}

func GetProductById(productId uint) (models.Product, error) {
	db := database.GetDB()
	product := models.Product{}
	err := db.First(&product, "id = ?", productId).Error
	return product, err
}

func DeleteProductById(productId uint) error {
	db := database.GetDB()
	product := models.Product{}
	product.ID = productId
	return db.Delete(&product).Error
}
