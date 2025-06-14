package controllers

import (
	"strconv"

	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
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

func (ph *GetUserByIdController) ValidateRequest(c *fiber.Ctx) (int, error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ph *GetUserByIdController) Run(c *fiber.Ctx) (err error) {
	id, err := ph.ValidateRequest(c)
	if err != nil {
		return ph.result.Error(c, err.Error())
	}
	user, err := ph.usecase.GetUserById(id)
	if err != nil {
		ph.result.Error(c, err)
		return err
	}

	ph.result.Ok(c, user, "User retrieved successfully")
	return nil
}
