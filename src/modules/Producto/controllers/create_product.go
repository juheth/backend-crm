package controllers

import (
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/usecases"
	"github.com/gofiber/fiber/v2"
)

type CreateProductController struct {
	usecase *usecases.CreateProduct
	result  *r.Result
}

func NewCreateProductController(uc *usecases.CreateProduct, r *r.Result) *CreateProductController {
	return &CreateProductController{usecase: uc, result: r}
}

func (c *CreateProductController) Run(ctx *fiber.Ctx) error {
	var req dto.CreateProductRequest

	if err := ctx.BodyParser(&req); err != nil {
		return c.result.Bad(ctx, "Error al parsear el body")
	}

	resp, err := c.usecase.Execute(req)
	if err != nil {
		return c.result.Error(ctx, err)
	}

	return c.result.Ok(ctx, resp)
}
