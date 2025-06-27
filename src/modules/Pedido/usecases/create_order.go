package usecases

import (
	"time"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/repository"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/utils"
)

type CreateOrder struct {
	repo repository.OrderRepository
}

func NewCreateOrder(repo *dao.MySQLOrderDao) *CreateOrder {
	return &CreateOrder{repo: repo}
}

func (uc *CreateOrder) Execute(input dto.CreateOrderRequest, userID int, clientDao *dao.MySQLClientDao) (*dto.OrderResponse, error) {
	if err := utils.ValidateCreateOrder(input, clientDao); err != nil {
		return nil, err
	}

	var total float64
	var items []entities.OrderItem
	var itemsResponse []dto.OrderItemResponse

	for _, item := range input.Items {
		price, err := uc.repo.GetProductPrice(item.ProductID)
		if err != nil {
			return nil, err
		}
		subtotal := price * float64(item.Quantity)
		total += subtotal

		items = append(items, entities.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: price,
			Subtotal:  subtotal,
		})

		itemsResponse = append(itemsResponse, dto.OrderItemResponse{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: price,
			Subtotal:  subtotal,
		})
	}

	order := &entities.Order{
		ClientID:        input.ClientID,
		UserID:          userID,
		Status:          "pendiente",
		Total:           total,
		OrderDate:       time.Now(),
		PaymentMethodID: input.PaymentMethodID,
		Items:           items,
	}

	if err := uc.repo.CreateOrder(order); err != nil {
		return nil, err
	}

	return &dto.OrderResponse{
		ID:        order.ID,
		ClientID:  order.ClientID,
		UserID:    order.UserID,
		Status:    order.Status,
		Total:     order.Total,
		OrderDate: order.OrderDate,
		Items:     itemsResponse,
	}, nil
}
