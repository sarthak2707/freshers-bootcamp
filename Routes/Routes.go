package Routes

import (
	"freshers-bootcamp/Controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetUsersEx1)
		grp1.POST("user", Controllers.CreateUserEx1)
		grp1.GET("user/:id", Controllers.GetUserByIDEx1)
		grp1.PUT("user/:id", Controllers.UpdateUserEx1)
		grp1.DELETE("user/:id", Controllers.DeleteUserEx1)
	}
	grp2 := r.Group("/student-api")
	{
		grp2.GET("student", Controllers.GetAllStudents)
		grp2.POST("student", Controllers.CreateStudent)
		grp2.GET("student/:id", Controllers.GetStudentById)
		grp2.PUT("student/:id", Controllers.UpdateStudent)
		grp2.GET("score/:id", Controllers.GetAllScoresForStudent)
		grp2.POST("score", Controllers.CreateTestScore)
		grp2.PUT("score/:id/:subject", Controllers.UpdateTestScore)
	}
	return r
}
