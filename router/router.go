package router

import (
	"challenge-2/controllers"
	"challenge-2/middlewares"

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
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/", middlewares.AdminAuthorization(), controllers.GetAllProducts)

		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), controllers.GetProductById)
		productRouter.DELETE("/:productId", middlewares.AdminAuthorization(), controllers.DeleteProductById)
	}
	return r
}
