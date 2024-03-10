package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type UpdateProductRequest struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

func (p *Product) TableName() string {
	return "products"
}
