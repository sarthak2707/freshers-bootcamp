package models

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique" binding:"required,email"`
	Phone     string `json:"phone" gorm:"unique" binding:"required,min=10,max=10"`
	Username  string `json:"username" gorm:"unique" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *Customer) TableName() string {
	return "customers"
}
