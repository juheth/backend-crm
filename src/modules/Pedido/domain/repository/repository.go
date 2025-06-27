package repository

import (
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/entities"
)

type OrderRepository interface {
	CreateOrder(order *entities.Order) error
	GetProductPrice(productID int) (float64, error)
}
