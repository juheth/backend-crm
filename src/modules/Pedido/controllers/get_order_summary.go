package controllers

import (
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Pedido/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetOrderSummaryController struct {
	usecase *usecases.GetOrderSummary
	result  *r.Result
}

func NewGetOrderSummaryController(uc *usecases.GetOrderSummary, r *r.Result) *GetOrderSummaryController {
	return &GetOrderSummaryController{
		usecase: uc,
		result:  r,
	}
}

func (ct *GetOrderSummaryController) Run(c *fiber.Ctx) error {
	summary, err := ct.usecase.Execute()
	if err != nil {
		return ct.result.Error(c, err)
	}
	return ct.result.Ok(c, summary)
}
