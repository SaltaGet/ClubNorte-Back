package middleware

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

func IsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*schemas.UserContext)
		if !user.IsAdmin {
			return c.Status(401).JSON(schemas.Response{
				Status:  false,
				Body:    nil,
				Message: "No tienes permiso para realizar esta accion",
			})
		}
		return c.Next()
	}
}