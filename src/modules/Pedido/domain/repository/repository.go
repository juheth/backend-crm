package repository

import (
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/entities"
)

type OrderRepository interface {
	CreateOrder(order *entities.Order) error
	GetProductPrice(productID int) (float64, error)
	GetAllOrders() ([]*entities.Order, error)
	GetOrderByID(id int) (*entities.Order, error)
	UpdateOrderStatus(orderID int, status string) error
	DeleteOrder(orderID int) error
	GetOrdersByClient(clientId int) ([]*entities.Order, error)
	GetOrdersByStatus(status string) ([]*entities.Order, error)
}
