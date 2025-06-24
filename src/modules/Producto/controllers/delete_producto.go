package controllers

import (
    "strconv"

    r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
    "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/usecases"
    "github.com/gofiber/fiber/v2"
)

type DeleteProductController struct {
    usecase *usecases.DeleteProduct
    result  *r.Result
}

func NewDeleteProductController(uc *usecases.DeleteProduct, r *r.Result) *DeleteProductController {
    return &DeleteProductController{usecase: uc, result: r}
}

func (c *DeleteProductController) Run(ctx *fiber.Ctx) error {
    idParam := ctx.Params("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.result.Bad(ctx, "ID inválido")
    }

    if err := c.usecase.Execute(id); err != nil {
        return c.result.Bad(ctx, err.Error())
    }

    return c.result.Ok(ctx, fiber.Map{"message": "Producto eliminado correctamente"})
}