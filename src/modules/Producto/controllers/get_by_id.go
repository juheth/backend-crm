package controllers

import (
	"strconv"

	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetProductByIDController struct {
	usecase *usecases.GetProductByID
	result  *r.Result
}

func NewGetProductByIDController(uc *usecases.GetProductByID, r *r.Result) *GetProductByIDController {
	return &GetProductByIDController{usecase: uc, result: r}
}

func (c *GetProductByIDController) Run(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return c.result.Bad(ctx, "ID de producto inválido")
	}
	product, err := c.usecase.Execute(id)
	if err != nil {
		return c.result.Bad(ctx, "Error al obtener el producto")
	}
	return c.result.Ok(ctx, fiber.StatusOK, product)
}
