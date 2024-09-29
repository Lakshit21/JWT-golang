package middleware

import (
	"server-backend/utils"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		// Validate the JWT token
		claims, err := utils.ValidateToken(tokenString[7:]) // Skip "Bearer " part
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		c.Locals("user_email", claims["email"])
		c.Locals("CLAIMs", claims)

		return c.Next()
	}
}
