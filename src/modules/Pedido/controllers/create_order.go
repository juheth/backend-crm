package controllers

import (
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	dao "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/infrastructure/db/dao"
	dto "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/domain/dto"

	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/usecases"
	utils "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/utils"
	"github.com/gofiber/fiber/v2"
)

type CreateOrderController struct {
	usecase        *usecases.CreateOrder
	result         *r.Result
	mysqlClientDao *dao.MySQLClientDao
}

func NewCreateOrderController(u *usecases.CreateOrder, r *r.Result, clienDao *dao.MySQLClientDao) *CreateOrderController {
	return &CreateOrderController{usecase: u, result: r, mysqlClientDao: clienDao}
}

func (c *CreateOrderController) Run(ctx *fiber.Ctx) error {
	var req dto.CreateOrderRequest
	if err := ctx.BodyParser(&req); err != nil {
		return c.result.Bad(ctx, "Error al parsear request")
	}

	userIDRaw := ctx.Locals("userID")
	if userIDRaw == nil {
		return c.result.Bad(ctx, "No autenticado")
	}
	userID, ok := userIDRaw.(int)
	if !ok {
		f, ok := userIDRaw.(float64)
		if ok {
			userID = int(f)
		} else {
			return c.result.Bad(ctx, "userID inválido")
		}
	}

	if err := utils.ValidateCreateOrder(req, c.mysqlClientDao); err != nil {
		return c.result.Bad(ctx, err.Error())
	}

	order, err := c.usecase.Execute(req, userID, c.mysqlClientDao)
	if err != nil {
		return c.result.Error(ctx, err)
	}

	return c.result.Ok(ctx, order)
}
