package middleware

import (
	"fmt"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

func IsAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*schemas.UserContext)
		if !user.IsAdmin {
			return schemas.HandleError(
				c,
				schemas.ErrorResponse(
					401,
					"No tienes permiso para realizar esta accion",
					fmt.Errorf("no tienes permiso para realizar esta accion rol requerido admin, rol de usuario %s", user.Role)))
		}
		return c.Next()
	}
}
