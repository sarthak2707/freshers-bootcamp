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
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
	}
	err = Models.GetStudentById(&student, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	c.BindJSON(&student)
	err = Models.UpdateStudent(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

func GetAllScoresForStudent(c *gin.Context) {
	var scores []Models.TestScore

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
	}

	err = Models.GetScoresByStudentId(&scores, uint(id))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, scores)
	}
}

func CreateTestScore(c *gin.Context) {
	var score Models.TestScore
	c.BindJSON(&score)
	err := Models.CreateTestScore(&score)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, score)
	}
}

func UpdateTestScore(c *gin.Context) {
	var score Models.TestScore

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
	}
	subject := c.Params.ByName("subject")

	err = Models.GetScoreForStudentBySubject(&score, uint(id), subject)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	fmt.Println(score)
	c.BindJSON(&score)
	fmt.Println(score)
	fmt.Println(subject)
	if score.Subject != subject {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}
	err = Models.UpdateTestScore(&score)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, score)
	}
}
