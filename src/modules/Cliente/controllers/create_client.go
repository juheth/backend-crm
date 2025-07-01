package controllers

import (
	authEntities "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/auth"
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/dto"
	clientEntities "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/domain/entities"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/usecases"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/Cliente/utils"
	"github.com/gofiber/fiber/v2"
)

type CreateClientController struct {
	usecase *usecases.CreateClient
	result  *r.Result
}

func NewCreateClientController(uc *usecases.CreateClient, r *r.Result) *CreateClientController {
	return &CreateClientController{
		usecase: uc,
		result:  r,
	}
}

func (ct *CreateClientController) Run(c *fiber.Ctx) error {
	var req dto.CreateClientRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return ct.result.Bad(c, "Error al parsear el body")
	}

	if err := utils.ValidateClientInput(req.Name, req.Email, req.Phone); err != nil {
		return ct.result.Bad(c, err.Error())
	}

	claims, ok := c.Locals("claims").(*authEntities.Claims)
	if !ok || claims == nil {
		return ct.result.Error(c, "Error al obtener los claims del token")
	}

	userID := claims.ID
	if userID <= 0 {
		return ct.result.Bad(c, "El ID del usuario creador es obligatorio y debe ser numérico")
	}

	client := &clientEntities.Client{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedBy: userID,
	}

	if err := ct.usecase.Execute(client); err != nil {
		return ct.result.Bad(c, err.Error())
	}

	return ct.result.Ok(c, client)
}
