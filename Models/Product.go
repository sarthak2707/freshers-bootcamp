package Models

import (
	"freshers-bootcamp/Config"
	_ "github.com/go-sql-driver/mysql"
)

func CreateProduct(product *Product) error {
	err := Config.DB.Create(product).Error
	if err != nil {
		return err
	}
	product.Message = "product successfully added"
	err = Config.DB.Save(product).Error
	return nil
}
