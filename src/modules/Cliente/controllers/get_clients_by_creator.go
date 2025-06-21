package controllers

import (
	"strconv"

	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetClientsByCreatorController struct {
	usecase *usecases.GetClientsByCreator
	result  *r.Result
}

func NewGetClientsByCreatorController(uc *usecases.GetClientsByCreator, r *r.Result) *GetClientsByCreatorController {
	return &GetClientsByCreatorController{
		usecase: uc,
		result:  r,
	}
}

func (ct *GetClientsByCreatorController) Run(c *fiber.Ctx) error {
	creatorIDStr := c.Params("creatorId")
	creatorID, err := strconv.Atoi(creatorIDStr)
	if err != nil || creatorID <= 0 {
		return ct.result.Bad(c, "ID del proveedor inválido")
	}

	clients, err := ct.usecase.Execute(creatorID)
	if err != nil {
		return ct.result.Error(c, err)
	}

	return ct.result.Ok(c, clients)
}
