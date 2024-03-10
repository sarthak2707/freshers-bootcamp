package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	CustomerID int    `json:"customer_id" binding:"required"`
	ProductID  int    `json:"product_id" binding:"required"`
	Quantity   int    `json:"quantity" binding:"required"`
	Amount     int    `json:"amount" binding:"required"`
	Status     string `json:"status" binding:"required" in:"order_placed,processed,failed" default:"order_placed"`
}

type OrderRequest struct {
	CustomerID int `json:"customer_id" binding:"required"`
	ProductID  int `json:"product_id" binding:"required"`
	Quantity   int `json:"quantity" binding:"required"`
}

func (o *Order) TableName() string {
	return "orders"
}
