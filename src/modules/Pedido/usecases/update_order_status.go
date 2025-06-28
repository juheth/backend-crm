package usecases

import (
	"errors"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/repository"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/utils"
	"gorm.io/gorm"
)

type UpdateOrderStatus struct {
	repo repository.OrderRepository
}

func NewUpdateOrderStatus(repo *dao.MySQLOrderDao) *UpdateOrderStatus {
	return &UpdateOrderStatus{repo: repo}
}

func (uc *UpdateOrderStatus) Execute(orderID int, status string) error {
	req := dto.UpdateOrderStatusRequest{Status: status}
	if err := utils.ValidateUpdateOrderStatus(req); err != nil {
		return err
	}

	order, err := uc.repo.GetOrderByID(orderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || order == nil {
			return errors.New("El pedido no existe")
		}
		return err
	}
	return uc.repo.UpdateOrderStatus(orderID, status)
}
