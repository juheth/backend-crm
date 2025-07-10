package controllers

import (
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/dto"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/usecases"
	"github.com/gofiber/fiber/v2"
)

type SendPasswordController struct {
	usecase *usecases.SendPassword
	result  *r.Result
}

func NewSendPasswordController(usecase *usecases.SendPassword, result *r.Result) *SendPasswordController {
	return &SendPasswordController{
		usecase: usecase,
		result:  result,
	}
}

func (spc *SendPasswordController) Run(c *fiber.Ctx) error {
	var input dto.SendPasswordRequest
	if err := c.BodyParser(&input); err != nil || input.Email == "" {
		return spc.result.Bad(c, "Correo inválido o faltante")
	}

	if err := spc.usecase.Execute(&input); err != nil {
		return spc.result.Error(c, err)
	}

	return spc.result.Ok(c, fiber.Map{
		"message": "Contraseña enviada correctamente al correo",
		"email":   input.Email,
	})
}
