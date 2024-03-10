package routes

import (
	"github.com/gin-gonic/gin"
	"manish/day4/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(gin.Logger(), gin.Recovery(), middlewares.AuthMiddleware())
	r.POST("/product", controllers.CreateProduct)
	r.PATCH("/product/:id", controllers.UpdateProduct)
	r.GET("/product/:id", controllers.GetProduct)
	r.GET("/products", controllers.GetAllProduct)

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.POST("/order", controllers.CreateOrder)
	r.GET("/order/:id", controllers.GetOrder)
	return r

}
