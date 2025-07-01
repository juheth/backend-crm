package usecases

import (
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"

	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/dto"
)

type GetOrderSummary struct {
	repo *dao.MySQLOrderDao
}

func NewGetOrderSummary(repo *dao.MySQLOrderDao) *GetOrderSummary {
	return &GetOrderSummary{repo: repo}
}

func (u *GetOrderSummary) Execute() (dto.OrderSummaryResponse, error) {
	totalPedidosMes, ventasTotalesMes, ticketPromedio, porEstado, err := u.repo.GetOrderSummary()
	if err != nil {
		return dto.OrderSummaryResponse{}, err
	}
	return dto.OrderSummaryResponse{
		TotalPedidosMes:  totalPedidosMes,
		VentasTotalesMes: ventasTotalesMes,
		TicketPromedio:   ticketPromedio,
		PorEstado:        porEstado,
	}, nil
}
