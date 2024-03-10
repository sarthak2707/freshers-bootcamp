package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"manish/day4/models"
	"manish/day4/repo"
	"time"
)

const (
	SecretKey    = "secret-key"
	RoleCustomer = "customer"
)

func createToken(username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	token.Valid = true

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, role string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return err
	}
	fmt.Println("claims", token.Claims, token.Valid, role)
	if !token.Valid || token.Claims.(jwt.MapClaims)["role"] != role {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func SignUp(user *models.Customer) error {

	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		user.Password = string(hashedPassword)
	}
	if err := repo.CreateCustomer(user); err != nil {
		return err
	}

	return nil
}

func Login(c *gin.Context, loginRequest *models.LoginRequest) error {
	user, err := repo.FindCustomerByUsername(loginRequest.Username)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return err
	}

	tokenString, err := createToken(loginRequest.Username, RoleCustomer)
	if err != nil {
		return err
	}

	c.Header("Authorization", tokenString)
	c.SetCookie("token", tokenString, 3600, "/", "", false, true)

	return nil
}
