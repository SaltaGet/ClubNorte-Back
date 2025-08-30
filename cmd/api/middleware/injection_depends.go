package middleware

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/dependencies"
	"github.com/gofiber/fiber/v2"
)

func InjectionDepends(deps *dependencies.MainContainer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("deps", deps)
		return c.Next()
	}
}