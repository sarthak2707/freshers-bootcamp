package Controllers

import (
	"fmt"
	"freshers-bootcamp/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type displayableProduct struct {
	Id          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
	Message     string `json:"message"`
}

func AddProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		output := displayableProduct{
			Id:          product.ID,
			ProductName: product.Name,
			Price:       product.Price,
			Quantity:    product.Quantity,
			Message:     product.Message,
		}
		c.JSON(200, output)
	}
}
