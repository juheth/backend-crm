package controllers

import (
    "strconv"
    r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
    "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/dto"
    "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/usecases"
    "github.com/gofiber/fiber/v2"
)

type UpdateOrderStatusController struct {
    usecase *usecases.UpdateOrderStatus
    result  *r.Result
}

func NewUpdateOrderStatusController(u *usecases.UpdateOrderStatus, r *r.Result) *UpdateOrderStatusController {
    return &UpdateOrderStatusController{usecase: u, result: r}
}

func (c *UpdateOrderStatusController) Run(ctx *fiber.Ctx) error {
    id, err := strconv.Atoi(ctx.Params("id"))
    if err != nil {
        return c.result.Bad(ctx, "ID inválido")
    }
    var req dto.UpdateOrderStatusRequest
    if err := ctx.BodyParser(&req); err != nil {
        return c.result.Bad(ctx, "Error al parsear request")
    }
    if err := c.usecase.Execute(id, req.Status); err != nil {
        return c.result.Bad(ctx, err.Error())
    }
    return c.result.Ok(ctx, "Estado actualizado correctamente")
}