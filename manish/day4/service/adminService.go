package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"manish/day4/models"
	"manish/day4/repo"
)

const (
	RoleAdmin = "admin"
)

func AdminSignUp(admin *models.Admin) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error while hashing password: ", err)
		return err
	}
	admin.Password = string(hashedPassword)
	if err = repo.CreateAdmin(admin); err != nil {
		fmt.Println("Error while creating admin: ", err)
		return err
	}
	return nil
}

func AdminLogin(c *gin.Context, loginRequest *models.LoginRequest) error {
	user, err := repo.FindAdminByUsername(loginRequest.Username)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return err
	}

	tokenString, err := createToken(loginRequest.Username, RoleAdmin)
	if err != nil {
		return err
	}

	c.Header("Authorization", tokenString)
	c.SetCookie("token", tokenString, 3600, "/", "", false, true)

	return nil
}
