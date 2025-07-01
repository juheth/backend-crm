package controllers

import (
	"fmt"

	config "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/config"
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/domain/dto"
	usecases "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/modules/User/usecases"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LoginUserController struct {
	usecase *usecases.LoginUser
	result  *r.Result
}

func NewLoginUserController(usecase *usecases.LoginUser, r *r.Result) *LoginUserController {
	return &LoginUserController{
		usecase: usecase,
		result:  r,
	}
}

func (ph *LoginUserController) ValidateRequest(c *fiber.Ctx) (dto.LoginDTO, error) {
	var user dto.LoginDTO

	if err := c.BodyParser(&user); err != nil {
		return user, fmt.Errorf(config.BAD_REQUEST, err.Error())
	}

	validate := validator.New()

	if err := validate.Struct(&user); err != nil {
		fieldErr := err.(validator.ValidationErrors)[0].Field()
		fieldTagErr := err.(validator.ValidationErrors)[0].Tag()

		if fieldErr == "Email" && fieldTagErr == "email" {
			return user, fmt.Errorf("El campo %s debe ser un email válido", fieldErr)
		}
		return user, fmt.Errorf("El campo %s es requerido", fieldErr)
	}

	if user.Password == "" {
		return user, fmt.Errorf("El campo password es requerido")
	}

	if len(user.Password) < 6 {
		return user, fmt.Errorf("La contraseña debe tener al menos 6 caracteres")
	}

	return user, nil
}

func (ph *LoginUserController) Run(c *fiber.Ctx) error {
	payload, err := ph.ValidateRequest(c)
	if err != nil {
		return ph.result.Bad(c, err.Error())
	}

	token, user, err := ph.usecase.Execute(payload)
	if err != nil {
		return ph.result.Bad(c, err.Error())
	}

	if user == nil {
		return ph.result.Bad(c, "Usuario no encontrado")
	}

	userDTO := dto.UserDTO{
		ID:   user.ID,
		Name: user.Name,
	}

	resp := dto.LoginResponse{
		User:  userDTO,
		Token: token,
	}

	return ph.result.Ok(c, resp)
}
