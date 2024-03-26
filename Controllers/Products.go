package Controllers

import (
	"fmt"
	"freshers-bootcamp/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, product.GetDisplayableProductWithMessageFromProduct())
	}
}

func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductById(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product.GetDisplayableProductFromProduct())
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product.GetDisplayableProductFromProduct())
	}
}

func GetAllProducts(c *gin.Context) {
	var products []Models.Product
	err := Models.GetAllProducts(&products)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, Models.GetDisplayableProductsFromProducts(products))
	}
}

func GetProductById(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductById(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(200, product.GetDisplayableProductFromProduct())
	}
}
