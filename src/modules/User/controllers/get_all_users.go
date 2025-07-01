package controllers

import (
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/dto"
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

func (ph *GetAllUsersController) Run(c *fiber.Ctx) error {
	users, err := ph.usecase.Execute()
	if err != nil {
		return ph.result.Error(c, err)
	}

	var usersResponse []dto.UserResponse
	for _, u := range users {
		usersResponse = append(usersResponse, dto.UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			Status:    u.Status,
			CreatedAt: u.CreatedAt,
		})
	}

	return ph.result.Ok(c, usersResponse)
}
