package controllers

import (
	"github.com/gin-gonic/gin"
	"manish/day4/models"
	"manish/day4/service"
	"strconv"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while unmarshalling product data",
			"message": err.Error(),
		})
		return
	}

	if err := service.CreateProduct(&product); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while creating product",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}

func UpdateProduct(c *gin.Context) {
	var (
		productUpdateRequest models.UpdateProductRequest
		product              *models.Product
		err                  error
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err = c.BindJSON(&productUpdateRequest); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while unmarshalling product_update_request data",
			"message": err.Error(),
		})
		return
	}

	if product, err = service.UpdateProduct(id, &productUpdateRequest); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while updating product",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Product updated successfully",
		"product": *product,
	})
}

func GetProduct(c *gin.Context) {
	var (
		product *models.Product
		err     error
	)
	id, err := strconv.Atoi(c.Param("id"))
	if product, err = service.GetProduct(id); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while getting product",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Product fetched successfully",
		"product": *product,
	})
}
func GetAllProduct(c *gin.Context) {
	var (
		products []models.Product
		err      error
	)
	if products, err = service.GetAllProduct(); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while getting products",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "Products fetched successfully",
		"products": products,
	})
}
