package repository

import (
	"challenge-2/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(product *models.Product) error
	GetProductById(id uint) *models.Product
	GetAllProduct() *[]models.Product
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) UpdateProduct(product *models.Product) error {
	return r.db.Model(product).Updates(models.Product{Title: product.Title, Description: product.Description}).Error
}

func (r *productRepository) DeleteProduct(product *models.Product) error {
	return r.db.Delete(product, "id = ?", product.ID).Error
}

func (r *productRepository) GetProductById(id uint) *models.Product {
	product := &models.Product{}
	err := r.db.First(product, "id = ?", id).Error
	if err != nil {
		return nil
	}
	return product
}

func (r *productRepository) GetAllProduct() *[]models.Product {
	products := &[]models.Product{}

	err := r.db.Find(&products).Error
	if err != nil {
		return nil
	}
	return products
}
