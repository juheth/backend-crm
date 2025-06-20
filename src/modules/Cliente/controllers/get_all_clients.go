package controllers

import (
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetAllClientsController struct {
	usecase *usecases.GetAllClients
	result  *r.Result
}

func NewGetAllClientsController(uc *usecases.GetAllClients, r *r.Result) *GetAllClientsController {
	return &GetAllClientsController{
		usecase: uc,
		result:  r,
	}
}

func (ct *GetAllClientsController) Run(c *fiber.Ctx) error {
	clients, err := ct.usecase.Execute()
	if err != nil {
		return ct.result.Bad(c, err.Error())
	}
	return ct.result.Ok(c, clients)
}
