package usecases

import (
	"time"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"

	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/repository"
)

type GetAllOrders struct {
	repo repository.OrderRepository
}

func NewGetAllOrders(repo *dao.MySQLOrderDao) *GetAllOrders {
	return &GetAllOrders{repo: repo}
}

func (uc *GetAllOrders) Execute(page, size int, status string, clientId int, dateFrom, dateTo *time.Time) ([]dto.OrderResponse, error) {
	orders, err := uc.repo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	var result []dto.OrderResponse
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
		result = append(result, dto.OrderResponse{
			ID:        order.ID,
			ClientID:  order.ClientID,
			UserID:    order.UserID,
			Status:    order.Status,
			Total:     order.Total,
			OrderDate: order.OrderDate,
			Items:     items,
		})
	}
	return result, nil
}
