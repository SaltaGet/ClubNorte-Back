package middleware

import (
	"slices"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

func AuthRoleMiddleware(roles []string) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*schemas.UserContext)

		if !slices.Contains(roles, user.Role) {
			return ctx.Status(403).JSON(schemas.Response{
				Status:  false,
				Body:    nil,
				Message: "No tienes permiso para realizar esta acción",
			})
		}

		return ctx.Next()
	}
}
