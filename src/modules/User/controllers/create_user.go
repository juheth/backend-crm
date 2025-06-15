package controllers

import (
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	usecases "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/usecases"
	"github.com/gofiber/fiber/v2"
)

type CreateUsersController struct {
	usecase *usecases.CreateUsers
	result  *r.Result
}

func NewCreateUsersController(usecase *usecases.CreateUsers, r *r.Result) *CreateUsersController {
	return &CreateUsersController{
		usecase: usecase,
		result:  r,
	}
}

func (ph *CreateUsersController) Run(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return ph.result.Bad(c, "Error al parsear el cuerpo de la solicitud")
	}

	if err := ph.usecase.Execute(&user); err != nil {
		return ph.result.Error(c, err)
	}

	if user.Name == "" || user.Email == "" {
		return ph.result.Bad(c, "El nombre y el correo electrónico son obligatorios")
	}

	return ph.result.Ok(c, "Usuario creado exitosamente")
}
