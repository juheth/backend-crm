package repository

import (
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/entities"
)

type ProductRepository interface {
	Create(product *entities.Product) error
	ExistsByName(name string) (bool, error)
	GetAll() ([]*entities.Product, error)
	GetProductByID(id int) (*entities.Product, error)
	UpdateProduct(product *entities.Product) error
	DeactivateProduct(id int) error
	DeleteProduct(id int) error
	ActivateProduct(id int) error
	GetProductByIDAnyStatus(id int) (*entities.Product, error)
	GetAllDeactivated() ([]*entities.Product, error)
	GetLowStock(threshold int) ([]*entities.Product, error)
}
