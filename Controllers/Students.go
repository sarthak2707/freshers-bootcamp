package Controllers

import (
	"fmt"
	"freshers-bootcamp/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllStudents(c *gin.Context) {
	var students []Models.Student
	err := Models.GetAllStudents(&students)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, students)
	}
}

func CreateStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	err := Models.CreateStudent(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

func GetStudentById(c *gin.Context) {
	var student Models.Student
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
	}
	err = Models.GetStudentById(&student, uint(id))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

func UpdateStudent(c *gin.Context) {
	var student Models.Student
	fmt.Println(student)
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
	}
	err = Models.GetStudentById(&student, uint(id))
	fmt.Println(student)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	fmt.Println(student)
	c.BindJSON(&student)
	fmt.Println(student)
	err = Models.UpdateStudent(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
