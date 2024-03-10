package repo

import (
	"manish/day4/config"
	"manish/day4/models"
)

func Create(product *models.Product) error {
	if err := config.DB.
		Table(product.TableName()).
		Create(&product).Error; err != nil {
		return err
	}

	return nil
}

func Update(product *models.Product) error {
	if err := config.DB.
		Table(product.TableName()).
		Where("id = ?", product.ID).
		Updates(product).Error; err != nil {
		return err
	}

	return nil
}

func FindProductByName(name string) (*models.Product, error) {
	var product models.Product

	if err := config.DB.
		Table("products").
		Where("name = ?", name).
		First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func FindProductById(id int) (*models.Product, error) {
	var product models.Product

	if err := config.DB.
		Table("products").
		Where("id = ?", id).
		First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func FindAllProduct() ([]models.Product, error) {
	var products []models.Product

	if err := config.DB.
		Table("products").
		Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func CreateCustomer(customer *models.Customer) error {
	if err := config.DB.
		Table(customer.TableName()).
		Create(&customer).Error; err != nil {
		return err
	}

	return nil
}

func FindCustomerByUsername(username string) (*models.Customer, error) {
	var customer models.Customer

	if err := config.DB.
		Table("customers").
		Where("username = ?", username).
		First(&customer).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func FindAdminByUsername(username string) (*models.Customer, error) {
	var customer models.Customer

	if err := config.DB.
		Table("admins").
		Where("username = ?", username).
		First(&customer).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func FindCustomerById(id int) (*models.Customer, error) {
	var customer models.Customer

	if err := config.DB.
		Table("customers").
		Where("id = ?", id).
		First(&customer).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

func CreateOrder(order *models.Order) error {
	if err := config.DB.
		Table(order.TableName()).
		Create(&order).Error; err != nil {
		return err
	}

	return nil
}

func FindOrderById(id string) (*models.Order, error) {
	var order models.Order

	if err := config.DB.
		Table("orders").
		Where("id = ?", id).
		First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func UpdateProductQuantity(id uint, quantity int) error {
	if err := config.DB.
		Table("products").
		Where("id = ?", id).
		Update("quantity", quantity).Error; err != nil {
		return err
	}

	return nil
}

func CreateAdmin(admin *models.Admin) error {
	if err := config.DB.
		Table("admins").
		Create(admin).Error; err != nil {
		return err
	}

	return nil
}

func FindLatestOrderByCustomerId(customerID int) (*models.Order, error) {
	var order models.Order

	if err := config.DB.
		Table("orders").
		Where("customer_id = ?", customerID).
		Where("status = ?", "order_placed").
		Order("created_at desc").
		First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
