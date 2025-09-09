package middleware

import (
	"fmt"

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
			return schemas.HandleError(c, schemas.ErrorResponse(401, "Se necesita auntenticación del punto de venta", fmt.Errorf("se necesita auntenticación del punto de venta")))
		}

		if user.Role == "admin" {
			return c.Next()
		}

		validatedPointSale, err := deps.AuthController.AuthService.AuthPointSale(user.ID, pointSale.ID)
		if err != nil {
			return schemas.HandleError(c, err)
		}

		c.Locals("point_sale", validatedPointSale)

		return c.Next()
	}
}
