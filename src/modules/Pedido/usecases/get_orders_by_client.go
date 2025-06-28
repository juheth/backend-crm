package usecases

import (
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"

	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/repository"
)

type GetOrdersByClient struct {
	repo repository.OrderRepository
}

func NewGetOrdersByClient(repo *dao.MySQLOrderDao) *GetOrdersByClient {
	return &GetOrdersByClient{repo: repo}
}

func (uc *GetOrdersByClient) Execute(clientId int) ([]dto.OrderResponse, error) {
	orders, err := uc.repo.GetOrdersByClient(clientId)
	if err != nil {
		return nil, err
	}
	var orderResponses []dto.OrderResponse
	for _, order := range orders {
		var items []dto.OrderItemResponse
		for _, item := range order.Items {
			items = append(items, dto.OrderItemResponse{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				UnitPrice: item.UnitPrice,
				Subtotal:  item.Subtotal,
			})
		}
		orderResponses = append(orderResponses, dto.OrderResponse{
			ID:        order.ID,
			ClientID:  order.ClientID,
			UserID:    order.UserID,
			Status:    order.Status,
			Total:     order.Total,
			OrderDate: order.OrderDate,
			Items:     items,
		})
	}
	return orderResponses, nil
}
