package controllers

import (
    "strconv"

    r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
    "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/usecases"
    "github.com/gofiber/fiber/v2"
)

type GetLowStockProductsController struct {
    usecase *usecases.GetLowStockProducts
    result  *r.Result
}

func NewGetLowStockProductsController(uc *usecases.GetLowStockProducts, r *r.Result) *GetLowStockProductsController {
    return &GetLowStockProductsController{usecase: uc, result: r}
}

func (c *GetLowStockProductsController) Run(ctx *fiber.Ctx) error {
    thresholdStr := ctx.Query("threshold", "5")
    threshold, err := strconv.Atoi(thresholdStr)
    if err != nil {
        return c.result.Bad(ctx, "Threshold inválido")
    }

    products, err := c.usecase.Execute(threshold)
    if err != nil {
        return c.result.Error(ctx, err.Error())
    }
    return c.result.Ok(ctx, products)
}