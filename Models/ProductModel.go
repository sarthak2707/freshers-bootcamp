package Models

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type Product struct {
	gorm.Model
	Name     string `json:"product_name"`
	Quantity uint   `json:"quantity"`
	Price    uint   `json:"price"`
	Message  string `json:"message"`
	mu       sync.Mutex
}

func (p *Product) TableName() string {
	return "products"
}
