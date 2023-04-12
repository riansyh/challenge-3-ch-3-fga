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

func CreateProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userID

	err := service.CreateProduct(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userID
	product.ID = uint(productId)

	err := service.UpdateProduct(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetAllProducts(c *gin.Context) {
	products, err := service.GetAllProducts()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))
	product, err := service.GetProductById(uint(productId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteProductById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	Product.UserID = userID
	Product.ID = uint(productId)

	err := service.DeleteProductById(uint(productId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product successfully deleted",
	})
}
