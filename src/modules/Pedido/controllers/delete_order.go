package controllers

import (
	"strconv"

	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/usecases"
	"github.com/gofiber/fiber/v2"
)

type DeleteOrderController struct {
	usecase *usecases.DeleteOrder
	result  *r.Result
}

func NewDeleteOrderController(u *usecases.DeleteOrder, r *r.Result) *DeleteOrderController {
	return &DeleteOrderController{usecase: u, result: r}
}

func (c *DeleteOrderController) Run(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return c.result.Bad(ctx, "ID inválido")
	}
	if err := c.usecase.Execute(id); err != nil {
		return c.result.Bad(ctx, err.Error())
	}
	return c.result.Ok(ctx, "Pedido eliminado correctamente")
}
