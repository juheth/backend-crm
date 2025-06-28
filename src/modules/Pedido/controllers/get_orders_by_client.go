package controllers

import (
	"strconv"

	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetOrdersByClientController struct {
	usecase *usecases.GetOrdersByClient
	result  *r.Result
}

func NewGetOrdersByClientController(u *usecases.GetOrdersByClient, r *r.Result) *GetOrdersByClientController {
	return &GetOrdersByClientController{usecase: u, result: r}
}

func (c *GetOrdersByClientController) Run(ctx *fiber.Ctx) error {
	clientId, err := strconv.Atoi(ctx.Params("clientId"))
	if err != nil {
		return c.result.Bad(ctx, "clientId inválido")
	}
	orders, err := c.usecase.Execute(clientId)
	if err != nil {
		return c.result.Error(ctx, err)
	}
	if len(orders) == 0 {
		return c.result.Ok(ctx, "El cliente no tiene pedidos")
	}
	return c.result.Ok(ctx, orders)
}
