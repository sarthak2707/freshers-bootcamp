package Routes

import (
	"freshers-bootcamp/Controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grpUser := r.Group("/user-api")
	{
		grpUser.GET("user", Controllers.GetUsersEx1)
		grpUser.POST("user", Controllers.CreateUserEx1)
		grpUser.GET("user/:id", Controllers.GetUserByIDEx1)
		grpUser.PUT("user/:id", Controllers.UpdateUserEx1)
		grpUser.DELETE("user/:id", Controllers.DeleteUserEx1)
	}
	grpStudent := r.Group("/student-api")
	{
		grpStudent.GET("student", Controllers.GetAllStudents)
		grpStudent.POST("student", Controllers.CreateStudent)
		grpStudent.GET("student/:id", Controllers.GetStudentById)
		grpStudent.PUT("student/:id", Controllers.UpdateStudent)
		grpStudent.GET("score/:id", Controllers.GetAllScoresForStudent)
		grpStudent.POST("score", Controllers.CreateTestScore)
		grpStudent.PUT("score/:id/:subject", Controllers.UpdateTestScore)
	}
	grpProduct := r.Group("/product-api")
	{
		grpProduct.POST("product", Controllers.AddProduct)
	}
	return r
}
