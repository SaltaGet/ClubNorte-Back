package middleware

import (
	"strings"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

// func BlockAccess() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		path := c.Path()

// 		if strings.HasPrefix(path, "/.") {
// 			return c.Status(fiber.StatusForbidden).JSON(schemas.Response{
// 				Status:  false,
// 				Body:    nil,
// 				Message: "Acceso denegado",
// 			})
// 		}

// 		return c.Next()
// 	}
// }
func BlockAccess() fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Path()

		// Bloquear acceso a archivos que empiecen con punto
		if strings.Contains(path, "/.") {
			return c.Status(fiber.StatusForbidden).JSON(schemas.Response{
				Status:  false,
				Body:    nil,
				Message: "Acceso denegado",
			})
		}

		return c.Next()
	}
}
