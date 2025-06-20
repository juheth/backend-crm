package infraestructure

import (
	"strings"

	common "dev.azure.com/proyects-crm/CRM-ECOMMERS/_git/Backend-crm/src/common/auth"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	parts := strings.SplitN(token, " ", 2) // Dividir solo en el primer espacio
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token should be in 'Bearer <token>' format"})
	}

	// Obtener solo el token
	token = parts[1]

	// Validar el token
	claims, err := common.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	// Almacenar los claims en el contexto
	c.Locals("claims", claims)

	// Continuar con la siguiente manejador
	return c.Next()
}
