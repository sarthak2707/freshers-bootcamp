package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"manish/day3/models"
	"manish/day3/service"
)

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": "Error while unmarshalling student data"})
		return
	}
	if err := service.CreateStudent(student); err != nil {
		c.JSON(400, gin.H{"error": "Error while creating student"})
		return
	}
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": "Error while unmarshalling student data"})
		return
	}
	if err := service.UpdateStudent(student); err != nil {
		c.JSON(400, gin.H{"error": "Error while updating student"})
		return
	}
}

func InsertMarks(c *gin.Context) {
	var insertMarksRequest models.InsertMarksRequest
	if err := c.BindJSON(&insertMarksRequest); err != nil {
		c.JSON(400, gin.H{"error": "Error while unmarshalling insert marks request"})
		return
	}

	var (
		student *models.Student
		err     error
	)
	if student, err = service.InsertMarks(insertMarksRequest); err != nil {
		c.JSON(400, gin.H{"error": "Error while inserting marks"})
		return
	}

	fmt.Println("Student: ", *student)
}
