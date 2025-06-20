package controllers

import (
	"strings"

	auth "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/auth"
	r "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/response"
	"github.com/gofiber/fiber/v2"
)

type RefreshTokenController struct {
	result *r.Result
}

func NewRefreshTokenController(result *r.Result) *RefreshTokenController {
	return &RefreshTokenController{result: result}
}

func (h *RefreshTokenController) Run(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return h.result.Bad(c, "Token no proporcionado")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return h.result.Bad(c, "Formato de token inválido")
	}

	tokenString := parts[1]

	newToken, userID, err := auth.RefreshToken(tokenString)
	if err != nil {
		return h.result.Bad(c, err.Error())
	}

	return h.result.Ok(c, fiber.Map{
		"user_id": userID,
		"token":   newToken,
	})
}
