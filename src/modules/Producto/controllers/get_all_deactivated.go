package controllers

import (
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetAllDeactivatedProductsController struct {
	usecase *usecases.GetAllDeactivatedProducts
	result  *r.Result
}

func NewGetAllDeactivatedProductsController(uc *usecases.GetAllDeactivatedProducts, r *r.Result) *GetAllDeactivatedProductsController {
	return &GetAllDeactivatedProductsController{usecase: uc, result: r}
}

func (c *GetAllDeactivatedProductsController) Run(ctx *fiber.Ctx) error {
	products, err := c.usecase.Execute()
	if err != nil {
		return c.result.Error(ctx, err.Error())
	}
	return c.result.Ok(ctx, products)
}
