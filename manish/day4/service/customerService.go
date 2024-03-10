package service

import (
	"golang.org/x/crypto/bcrypt"
	"manish/day4/models"
	"manish/day4/repo"
)

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

func Login(loginRequest *models.LoginRequest) error {
	user, err := repo.FindCustomerByUsername(loginRequest.Username)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return err
	}

	return nil
}
