package middleware

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/logging"
	"github.com/DanielChachagua/Club-Norte-Back/internal/dependencies"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/gofiber/fiber/v2"
)

func AuthPointSaleMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		deps := c.Locals("deps").(*dependencies.MainContainer)
		user := c.Locals("user").(*schemas.UserContext)
		pointSale, ok := c.Locals("point_sale").(*schemas.PointSaleContext)

		if !ok || pointSale == nil {
			return c.Status(401).JSON(schemas.Response{
				Status:  false,
				Body:    nil,
				Message: "Se necesita autenticacion al punto de venta",
			})
		}

		if user.Role == "admin" {
			return c.Next()
		}

		validatedPointSale, err := deps.AuthController.AuthService.AuthPointSale(user.ID, pointSale.ID)
		if err != nil {
			logging.ERROR("Error: %s", err.Error())
			return c.Status(fiber.StatusUnauthorized).JSON(schemas.Response{
				Status:  false,
				Body:    nil,
				Message: "Error al obtener el punto de venta",
			})
		}

		c.Locals("point_sale", validatedPointSale)

		return c.Next()
	}
}
