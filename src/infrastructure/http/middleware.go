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

	parts := strings.SplitN(token, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token should be in 'Bearer <token>' format"})
	}

	token = parts[1]

	claims, err := common.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	if claims != nil {
		c.Locals("userID", claims.ID)
	}
	c.Locals("claims", claims)

	return c.Next()
}
