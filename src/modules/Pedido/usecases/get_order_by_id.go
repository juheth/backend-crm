package usecases

import (
	"errors"

	"gorm.io/gorm"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/repository"
)

type GetOrderByID struct {
	repo repository.OrderRepository
}

func NewGetOrderByID(repo *dao.MySQLOrderDao) *GetOrderByID {
	return &GetOrderByID{repo: repo}
}

func (uc *GetOrderByID) Execute(id int) (*dto.OrderResponse, error) {
	order, err := uc.repo.GetOrderByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Pedido no encontrado")
		}
		return nil, err
	}
	if order == nil {
		return nil, errors.New("Pedido no encontrado")
	}

	var itemsResponse []dto.OrderItemResponse
	for _, item := range order.Items {
		itemsResponse = append(itemsResponse, dto.OrderItemResponse{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
			Subtotal:  item.Subtotal,
		})
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
