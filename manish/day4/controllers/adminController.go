package controllers

import (
	"github.com/gin-gonic/gin"
	"manish/day4/models"
	"manish/day4/service"
)

func AdminSignUp(c *gin.Context) {
	var user models.Admin
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while unmarshalling user data",
			"message": err.Error(),
		})
		return
	}

	if err := service.AdminSignUp(&user); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while creating user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Admin created successfully",
		"user":    user,
	})
}

func AdminLogin(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while unmarshalling login request data",
			"message": err.Error(),
		})
		return
	}

	if err := service.AdminLogin(c, &loginRequest); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while login",
			"message": "Please check username/password and try again.",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Admin login successfully",
	})
}
