package usecases

import (
	"errors"

	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/repository"
)

type DeleteOrder struct {
	repo repository.OrderRepository
}

func NewDeleteOrder(repo *dao.MySQLOrderDao) *DeleteOrder {
	return &DeleteOrder{repo: repo}
}

func (uc *DeleteOrder) Execute(orderID int) error {
	order, err := uc.repo.GetOrderByID(orderID)
	if err != nil || order == nil {
		return errors.New("El pedido no existe")
	}
	if order.Status == "pagado" || order.Status == "entregado" {
		return errors.New("No se puede eliminar un pedido pagado o entregado")
	}
	return uc.repo.DeleteOrder(orderID)
}
