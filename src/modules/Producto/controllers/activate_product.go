package controllers

import (
    "strconv"

    r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
    "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/usecases"
    "github.com/gofiber/fiber/v2"
)

type ActivateProductController struct {
    usecase *usecases.ActivateProduct
    result  *r.Result
}

func NewActivateProductController(uc *usecases.ActivateProduct, r *r.Result) *ActivateProductController {
    return &ActivateProductController{usecase: uc, result: r}
}

func (c *ActivateProductController) Run(ctx *fiber.Ctx) error {
    idParam := ctx.Params("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.result.Bad(ctx, "ID inválido")
    }

    product, err := c.usecase.Execute(id)
    if err != nil {
        return c.result.Bad(ctx, err.Error())
    }

    return c.result.Ok(ctx, product)
}