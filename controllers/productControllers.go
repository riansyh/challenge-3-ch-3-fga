package controllers

import (
	"challenge-2/helpers"
	"challenge-2/models"
	"challenge-2/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	GetProductById(c *gin.Context)
	GetAllProduct(c *gin.Context)
}

type productController struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) *productController {
	return &productController{service: service}
}

func (s *productController) CreateProduct(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.MustGet("userData").(jwt.MapClaims)["id"].(float64)

	Product := models.Product{}
	Product.UserID = uint(userID)

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	err := s.service.CreateProduct(&Product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func (s *productController) UpdateProduct(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.MustGet("userData").(jwt.MapClaims)["id"].(float64)

	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))

	Product.UserID = uint(userID)
	Product.ID = uint(productId)

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	err := s.service.UpdateProduct(&Product, uint(userID), uint(productId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func (s *productController) GetProductById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	Product.UserID = userID
	Product.ID = uint(productId)

	socmed, err := s.service.GetProductById(Product.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, socmed)
}

func (s *productController) DeleteProduct(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userID := c.MustGet("userData").(jwt.MapClaims)["id"].(float64)

	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))

	Product.UserID = uint(userID)
	Product.ID = uint(productId)

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	err := s.service.DeleteProduct(&Product, uint(userID), uint(productId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product successfully deleted",
	})
}

func (s *productController) GetAllProduct(c *gin.Context) {
	result, err := s.service.GetAllProduct()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
