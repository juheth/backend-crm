package controllers

import (
	"strconv"

	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetClientByIDController struct {
	usecase *usecases.GetClientByID
	result  *r.Result
}

func NewGetClientByIDController(uc *usecases.GetClientByID, r *r.Result) *GetClientByIDController {
	return &GetClientByIDController{
		usecase: uc,
		result:  r,
	}
}

func (ct *GetClientByIDController) Run(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ct.result.Bad(c, "ID inválido")
	}
	client, err := ct.usecase.Execute(id)
	if err != nil {
		return ct.result.Bad(c, err.Error())
	}
	return ct.result.Ok(c, client)
}
