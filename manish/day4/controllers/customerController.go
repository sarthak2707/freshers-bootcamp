package controllers

import (
	"github.com/gin-gonic/gin"
	"manish/day4/models"
	"manish/day4/service"
)

func SignUp(c *gin.Context) {
	var user models.Customer
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while unmarshalling user data",
			"message": err.Error(),
		})
		return
	}

	if err := service.SignUp(&user); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while creating user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func Login(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while unmarshalling login request data",
			"message": err.Error(),
		})
		return
	}

	if err := service.Login(&loginRequest); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while login",
			"message": "Please check username/password and try again.",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User login successfully",
	})
}
