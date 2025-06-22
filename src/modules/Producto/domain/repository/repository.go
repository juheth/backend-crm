package repository

import (
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/entities"
)

type ProductRepository interface {
	Create(product *entities.Product) error
	ExistsByName(name string) (bool, error)
}
