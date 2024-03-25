package Controllers

import (
	"fmt"
	"freshers-bootcamp/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUsersEx1 ... Get all users
func GetUsersEx1(c *gin.Context) {
	var user []Models.UserEx1
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUserEx1 ... Create UserEx1
func CreateUserEx1(c *gin.Context) {
	var user Models.UserEx1
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUserByIDEx1 ... Get the user by id
func GetUserByIDEx1(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.UserEx1
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateUserEx1 ... Update the user information
func UpdateUserEx1(c *gin.Context) {
	var user Models.UserEx1
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// DeleteUserEx1 ... Delete the user
func DeleteUserEx1(c *gin.Context) {
	var user Models.UserEx1
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
