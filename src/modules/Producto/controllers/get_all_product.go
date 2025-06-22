package controllers

import (
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetAllProductsController struct {
	usecase *usecases.GetAllProducts
	result  *r.Result
}

func NewGetAllProductsController(uc *usecases.GetAllProducts, r *r.Result) *GetAllProductsController {
	return &GetAllProductsController{usecase: uc, result: r}
}

func (c *GetAllProductsController) Run(ctx *fiber.Ctx) error {
	products, err := c.usecase.Execute()
	if err != nil {
		return c.result.Error(ctx, err.Error())
	}
	return c.result.Ok(ctx, products)
}
