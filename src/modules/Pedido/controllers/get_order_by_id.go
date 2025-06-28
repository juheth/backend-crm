package controllers

import (
	"strconv"

	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetOrderByIDController struct {
	usecase *usecases.GetOrderByID
	result  *r.Result
}

func NewGetOrderByIDController(u *usecases.GetOrderByID, r *r.Result) *GetOrderByIDController {
	return &GetOrderByIDController{usecase: u, result: r}
}

func (c *GetOrderByIDController) Run(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return c.result.Bad(ctx, "ID inválido")
	}
	order, err := c.usecase.Execute(id)
	if err != nil {
		return c.result.Error(ctx, err)
	}
	return c.result.Ok(ctx, order)
}
