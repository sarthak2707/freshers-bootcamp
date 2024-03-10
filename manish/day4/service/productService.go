package service

import (
	"github.com/ulule/deepcopier"
	"manish/day4/models"
	"manish/day4/repo"
	"sync"
)

var (
	mutex sync.Mutex
)

func CreateProduct(product *models.Product) error {
	mutex.Lock()
	defer mutex.Unlock()
	if product1, err := repo.FindProductByName(product.Name); product1 != nil {
		product1.Quantity += product.Quantity

		if err = repo.Update(product1); err != nil {
			return err
		}

		*product = *product1
		return nil
	}
	if err := repo.Create(product); err != nil {
		return err
	}

	return nil
}

func UpdateProduct(id int, product *models.UpdateProductRequest) (*models.Product, error) {
	mutex.Lock()
	defer mutex.Unlock()
	product1, err := repo.FindProductById(id)
	if err != nil {
		return nil, err
	}

	deepcopier.Copy(product).To(product1)
	if err = repo.Update(product1); err != nil {
		return nil, err
	}

	return product1, nil
}

func GetProduct(id int) (*models.Product, error) {
	product, err := repo.FindProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func GetAllProduct() ([]models.Product, error) {
	products, err := repo.FindAllProduct()
	if err != nil {
		return nil, err
	}

	return products, nil
}
