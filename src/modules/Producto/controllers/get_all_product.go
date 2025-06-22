package controllers

import (
	"github.com/gofiber/fiber/v2"
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	usecases "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/usecases"
)

type GetAllProductsController struct {
	usecase *usecases.GetAllProducts
	result  *r.Result
}

func NewGetAllProductsController(usecase *usecases.GetAllProducts, result *r.Result) *GetAllProductsController {
	return &GetAllProductsController{usecase, result}
}

func (ph *GetAllProductsController) Run(c *fiber.Ctx) error {
	products, err := ph.usecase.Execute()
	if err != nil {
		return ph.result.Error(c, err)
	}
	return ph.result.Ok(c, products)
}
