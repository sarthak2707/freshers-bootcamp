package controllers

import (
	"github.com/gin-gonic/gin"
	"manish/day4/models"
	"manish/day4/service"
	"time"
)

func CreateOrder(c *gin.Context) {
	var (
		orderRequest models.OrderRequest
		order        *models.Order
		err          error
	)

	if err = c.BindJSON(&orderRequest); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while unmarshalling order data",
			"message": err.Error(),
		})
		return
	}

	if order, err = service.FetchLatestOrderForCustomer(orderRequest.CustomerID); err != nil {
		c.JSON(400, gin.H{
			"error": "Error while fetching latest order for customer",
		})
	}

	if time.Now().Unix()-order.CreatedAt.Unix() < int64(time.Minute.Minutes()*5) {
		c.JSON(400, gin.H{
			"error":   "Order already placed in last minute",
			"message": "Please try again after a minute",
		})
		return
	}

	if order, err = service.CreateOrder(&orderRequest); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while creating order",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Order created successfully",
		"order":   *order,
	})

}

func GetOrder(c *gin.Context) {
	var (
		order *models.Order
		err   error
	)
	id := c.Param("id")
	if order, err = service.GetOrder(id); err != nil {
		c.JSON(400, gin.H{
			"error":   "Error while fetching order",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"order": *order,
	})
}
