package routes

import (
	"github.com/gin-gonic/gin"
	"manish/day3/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/create_student", controllers.CreateStudent)
	r.PATCH("/update_student", controllers.UpdateStudent)
	r.POST("/insert_marks", controllers.InsertMarks)

	return r
}
