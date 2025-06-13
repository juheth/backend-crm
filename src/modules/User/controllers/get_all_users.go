package controllers

import (
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	usecases "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetAllUsersController struct {
	usecase *usecases.GetAllUsers
	result  *r.Result
}

func NewGetAllUsersController(usecase *usecases.GetAllUsers, r *r.Result) *GetAllUsersController {
	return &GetAllUsersController{
		usecase: usecase,
		result:  r,
	}
}

func (ph *GetAllUsersController) Run(c *fiber.Ctx) (err error) {
	var user entities.User

	ph.result.Ok(c, user)
	return
}
