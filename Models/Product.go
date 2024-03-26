package Models

import (
	"fmt"
	"freshers-bootcamp/Config"
	_ "github.com/go-sql-driver/mysql"
)

type DisplayableProductWithMessage struct {
	Id          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
	Message     string `json:"message"`
}

type DisplayableProduct struct {
	Id          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
}

func (product *Product) GetDisplayableProductWithMessageFromProduct() (dp DisplayableProductWithMessage) {
	dp = DisplayableProductWithMessage{
		Id:          product.ID,
		ProductName: product.Name,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Message:     product.Message,
	}
	return dp
}

func (product *Product) GetDisplayableProductFromProduct() (dp DisplayableProduct) {
	dp = DisplayableProduct{
		Id:          product.ID,
		ProductName: product.Name,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}
	return dp
}

func GetDisplayableProductsFromProducts(products []Product) (dps []DisplayableProduct) {
	for _, product := range products {
		op := product.GetDisplayableProductFromProduct()
		dps = append(dps, op)
	}
	return dps
}

func CreateProduct(product *Product) error {
	err := Config.DB.Create(product).Error
	if err != nil {
		return err
	}
	product.Message = "product successfully added"
	err = Config.DB.Save(product).Error
	return nil
}

func UpdateProduct(product *Product) (err error) {
	fmt.Println(product)
	if err := Config.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductById(product *Product, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProducts(products *[]Product) (err error) {
	if err := Config.DB.Find(products).Error; err != nil {
		return err
	}
	return nil
}
