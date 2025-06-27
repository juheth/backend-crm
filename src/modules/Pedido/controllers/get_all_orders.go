package controllers

import (
    "strconv"
    "time"

    r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
    "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/usecases"
    "github.com/gofiber/fiber/v2"
)

type GetAllOrdersController struct {
    usecase *usecases.GetAllOrders
    result  *r.Result
}

func NewGetAllOrdersController(u *usecases.GetAllOrders, r *r.Result) *GetAllOrdersController {
    return &GetAllOrdersController{usecase: u, result: r}
}

func (c *GetAllOrdersController) Run(ctx *fiber.Ctx) error {
    page, _ := strconv.Atoi(ctx.Query("page", "1"))
    size, _ := strconv.Atoi(ctx.Query("size", "10"))
    status := ctx.Query("status")
    clientId, _ := strconv.Atoi(ctx.Query("clientId", "0"))
    dateFrom := ctx.Query("dateFrom")
    dateTo := ctx.Query("dateTo")

    var from, to *time.Time
    if dateFrom != "" {
        t, err := time.Parse("2006-01-02", dateFrom)
        if err == nil {
            from = &t
        }
    }
    if dateTo != "" {
        t, err := time.Parse("2006-01-02", dateTo)
        if err == nil {
            to = &t
        }
    }

    orders, err := c.usecase.Execute(page, size, status, clientId, from, to)
    if err != nil {
        return c.result.Error(ctx, err)
    }
    return c.result.Ok(ctx, orders)
}