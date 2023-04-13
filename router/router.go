package router

import (
	"challenge-2/controllers"
	"challenge-2/database"
	"challenge-2/middlewares"
	"challenge-2/repository"
	"challenge-2/service"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRepository := repository.NewProductRepository(database.GetDB())
		productService := service.NewProductService(productRepository)
		productController := controllers.NewProductController(productService)

		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", productController.CreateProduct)
		productRouter.GET("/", middlewares.AdminAuthorization(), productController.GetAllProduct)

		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), productController.UpdateProduct)
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), productController.GetProductById)
		productRouter.DELETE("/:productId", middlewares.AdminAuthorization(), productController.DeleteProduct)
	}
	return r
}
