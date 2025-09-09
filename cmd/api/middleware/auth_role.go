package middleware

import (
	"fmt"
	"slices"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

func AuthRoleMiddleware(roles []string) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*schemas.UserContext)

		if !slices.Contains(roles, user.Role) {
			return schemas.HandleError(ctx, schemas.ErrorResponse(403, "No tienes permiso para realizar estaacción", fmt.Errorf("no tienes permiso para realizar esta acción")))
		}

		return ctx.Next()
	}
}
