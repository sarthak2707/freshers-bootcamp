package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"manish/day4/config"
	"manish/day4/models"
	"manish/day4/repo"
)

func CreateOrder(request *models.OrderRequest) (*models.Order, error) {
	var (
		product *models.Product
		err     error
		order   models.Order
	)
	mutex.Lock()
	defer mutex.Unlock()
	err = config.DB.Transaction(func(tx *gorm.DB) error {
		if product, err = repo.FindProductById(request.ProductID); err != nil {
			fmt.Println("Error while finding product by id: ", err)
			return err
		}

		if _, err = repo.FindCustomerById(request.CustomerID); err != nil {
			fmt.Println("Error while finding customer by id: ", err)
			return err
		}

		if product.Quantity < request.Quantity {
			fmt.Println("Product out of stock. Quantity available: ", product.Quantity, " Quantity requested: ", request.Quantity)
			return fmt.Errorf("product out of stock")
		}
		order = models.Order{
			CustomerID: request.CustomerID,
			ProductID:  request.ProductID,
			Quantity:   request.Quantity,
			Amount:     request.Quantity * product.Price,
			Status:     "order_placed",
		}

		if err = repo.CreateOrder(&order); err != nil {
			fmt.Println("Error while creating order: ", err)
			return err
		}

		if err = repo.UpdateProductQuantity(product.ID, product.Quantity-request.Quantity); err != nil {
			fmt.Println("Error while updating product quantity: ", err)
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &order, nil
}

func GetOrder(id string) (*models.Order, error) {
	order, err := repo.FindOrderById(id)
	if err != nil {
		fmt.Println("Error while fetching order by id: ", err)
		return nil, err
	}

	return order, nil
}
