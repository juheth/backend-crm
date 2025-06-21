package controllers

import (
	"strconv"

	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/usecases"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/utils"
	"github.com/gofiber/fiber/v2"
)

type UpdateClientController struct {
	usecase *usecases.UpdateClient
	result  *r.Result
}

func NewUpdateClientController(uc *usecases.UpdateClient, r *r.Result) *UpdateClientController {
	return &UpdateClientController{
		usecase: uc,
		result:  r,
	}
}

func (ct *UpdateClientController) Run(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return ct.result.Bad(c, "ID inválido")
	}

	var req dto.UpdateClientRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return ct.result.Bad(c, "Error al parsear el body")
	}

	if err := utils.ValidateClientInput(req.Name, req.Email, req.Phone); err != nil {
		return ct.result.Bad(c, err.Error())
	}

	client := &entities.Client{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}

	if err := ct.usecase.Execute(client); err != nil {
		return ct.result.Bad(c, err.Error())
	}

	return ct.result.Ok(c, client)
}
