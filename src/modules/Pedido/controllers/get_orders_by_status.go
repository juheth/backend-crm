package controllers

import (
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetOrdersByStatusController struct {
	usecase *usecases.GetOrdersByStatus
	result  *r.Result
}

func NewGetOrdersByStatusController(uc *usecases.GetOrdersByStatus, r *r.Result) *GetOrdersByStatusController {
	return &GetOrdersByStatusController{
		usecase: uc,
		result:  r,
	}
}

func (c *GetOrdersByStatusController) Run(ctx *fiber.Ctx) error {
	status := ctx.Query("status")
	if status == "" {
		return c.result.Bad(ctx, "El parámetro status es obligatorio")
	}
	orders, err := c.usecase.Execute(status)
	if err != nil {
		return c.result.Error(ctx, err)
	}
	if len(orders) == 0 {
		return c.result.Ok(ctx, "No hay pedidos con ese estado")
	}
	return c.result.Ok(ctx, orders)
}
