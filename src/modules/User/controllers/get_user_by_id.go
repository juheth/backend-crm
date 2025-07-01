package controllers

import (
	"strconv"

	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/dto"
	usecases "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/usecases"
	"github.com/gofiber/fiber/v2"
)

type GetUserByIdController struct {
	usecase *usecases.GetUserById
	result  *r.Result
}

func NewGetUserByIdController(usecase *usecases.GetUserById, r *r.Result) *GetUserByIdController {
	return &GetUserByIdController{
		usecase: usecase,
		result:  r,
	}
}

func (ph *GetUserByIdController) validateRequest(c *fiber.Ctx) (int, error) {
	return strconv.Atoi(c.Params("id"))
}

func (ph *GetUserByIdController) Run(c *fiber.Ctx) (err error) {
	id, err := ph.validateRequest(c)
	if err != nil {
		return ph.result.Bad(c, err.Error())
	}

	user, err := ph.usecase.Execute(id)

	if user.ID == 0 {
		return ph.result.Bad(c, "Usuario no encontrado")
	}

	if err != nil {
		ph.result.Error(c, err)
		return err
	}

	userResponse := dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
	}

	ph.result.Ok(c, userResponse)
	return nil
}
