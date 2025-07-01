package controllers

import (
	"strconv"

	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	usecases "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/usecases"
	utils "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/utils" // Agrega este import
	"github.com/gofiber/fiber/v2"
)

type UpdateUserController struct {
	usecase *usecases.UpdateUser
	result  *r.Result
}

func NewUpdateUserController(usecase *usecases.UpdateUser, r *r.Result) *UpdateUserController {
	return &UpdateUserController{
		usecase: usecase,
		result:  r,
	}
}

func (ph *UpdateUserController) Run(c *fiber.Ctx) error {
	var user entities.User

	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return ph.result.Bad(c, "ID inválido")
	}

	if err := c.BodyParser(&user); err != nil {
		return ph.result.Bad(c, "Error al parsear el body: "+err.Error())
	}

	user.ID = id

	if user.Name == "" {
		return ph.result.Bad(c, "El nombre es obligatorio")
	}
	if user.Email == "" {
		return ph.result.Bad(c, "El correo electrónico es obligatorio")
	}

	if user.Password != "" {
		hashed, err := utils.HashPassword(user.Password)
		if err != nil {
			return ph.result.Bad(c, "Error al cifrar la contraseña")
		}
		user.Password = hashed
	}

	if err := ph.usecase.Execute(&user); err != nil {
		return ph.result.Error(c, err)
	}

	return ph.result.Ok(c, "Usuario actualizado correctamente")
}
