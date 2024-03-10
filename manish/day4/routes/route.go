package routes

import (
	"github.com/gin-gonic/gin"
	"manish/day4/controllers"
	"manish/day4/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	adminRoutes := r.Group("/admin")
	userRoutes := r.Group("/")

	adminRoutes.POST("/signup", controllers.AdminSignUp)
	adminRoutes.POST("/login", controllers.AdminLogin)

	userRoutes.POST("/signup", controllers.SignUp)
	userRoutes.POST("/login", controllers.Login)
	userRoutes.GET("/logout", controllers.Logout)

	adminRoutes.Use(middlewares.AuthMiddleware("admin"))
	userRoutes.Use(middlewares.AuthMiddleware("customer"))

	adminRoutes.POST("/product", controllers.CreateProduct).Use(middlewares.AuthMiddleware("admin"))
	adminRoutes.PATCH("/product/:id", controllers.UpdateProduct).Use(middlewares.AuthMiddleware("admin"))
	adminRoutes.GET("/product/:id", controllers.GetProduct).Use(middlewares.AuthMiddleware("admin"))
	adminRoutes.GET("/products", controllers.GetAllProduct).Use(middlewares.AuthMiddleware("admin"))

	userRoutes.GET("/product/:id", controllers.GetProduct).Use(middlewares.AuthMiddleware("customer"))
	userRoutes.GET("/products", controllers.GetAllProduct).Use(middlewares.AuthMiddleware("customer"))

	userRoutes.POST("/order", controllers.CreateOrder).Use(middlewares.AuthMiddleware("customer"))
	userRoutes.GET("/order/:id", controllers.GetOrder).Use(middlewares.AuthMiddleware("customer"))

	return r
}
