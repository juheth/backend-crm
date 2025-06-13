package controllers

import (
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/entities"
	usecases "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/usecases"
	"github.com/gofiber/fiber/v2"
)

type FindAllUserController struct {
	usecase *usecases.FindAllUser
	result  *r.Result
}

func NewFindAllUserController(usecase *usecases.FindAllUser, r *r.Result) *FindAllUserController {
	return &FindAllUserController{
		usecase: usecase,
		result:  r,
	}
}

func (ph *FindAllUserController) Run(c *fiber.Ctx) (err error) {
	var user entities.User

	ph.result.Ok(c, user)
	return
}
