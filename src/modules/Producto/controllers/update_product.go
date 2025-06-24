package controllers

import (
    "strconv"

    r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
    "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/domain/dto"
    "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Producto/usecases"
    "github.com/gofiber/fiber/v2"
)

type UpdateProductController struct {
    usecase *usecases.UpdateProduct
    result  *r.Result
}

func NewUpdateProductController(uc *usecases.UpdateProduct, r *r.Result) *UpdateProductController {
    return &UpdateProductController{usecase: uc, result: r}
}

func (c *UpdateProductController) Run(ctx *fiber.Ctx) error {
    idParam := ctx.Params("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.result.Bad(ctx, "ID inválido")
    }

    var req dto.UpdateProductRequest
    if err := ctx.BodyParser(&req); err != nil {
        return c.result.Bad(ctx, "Error al parsear el body")
    }

    resp, err := c.usecase.Execute(id, req)
    if err != nil {
        return c.result.Bad(ctx, err.Error())
    }

    return c.result.Ok(ctx, resp)
}