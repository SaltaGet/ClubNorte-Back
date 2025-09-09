package middleware

import (
	"fmt"

	"github.com/DanielChachagua/Club-Norte-Back/internal/dependencies"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/DanielChachagua/Club-Norte-Back/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("access_token") // Obtenemos la cookie
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(schemas.Response{
				Status:  false,
				Body:    nil,
				Message: "No autorizado",
			})
		}

		claims, err := utils.VerifyToken(token)
		if err != nil {
			return schemas.HandleError(c, schemas.ErrorResponse(401, "Token inválido", err))
		}

		mapClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			return schemas.HandleError(c, schemas.ErrorResponse(401, "Claims inválidos", err))
		}

		email := getStringClaim(mapClaims, "email")
		if email == "" {
			return schemas.HandleError(c, schemas.ErrorResponse(401, "Email inválido", fmt.Errorf("email inválido")))
		}

		deps, ok := c.Locals("deps").(*dependencies.MainContainer)
		if !ok {
			return schemas.HandleError(c, fmt.Errorf("error al obtener dependencias"))
		}

		user, err := deps.AuthController.AuthService.AuthUser(email)
		if err != nil {
			return schemas.HandleError(c, err)
		}

		if !user.IsActive {
			return schemas.HandleError(c, schemas.ErrorResponse(403, "Usuario no activo", fmt.Errorf("usuario no activo")))
		}

		pointSaleID := getIntClaim(mapClaims, "point_sale_id")
		pointSale := getStringClaim(mapClaims, "point_sale")

		c.Locals("user", user)

		if pointSaleID == -1 || pointSale == "" {
			c.Locals("point_sale", nil)
		} else {
			c.Locals("point_sale", &schemas.PointSaleContext{
				ID:   uint(pointSaleID),
				Name: pointSale,
			})
		}

		if user.IsAdmin {
			return c.Next()
		}

		return c.Next()
	}
}

func getBoolClaim(claims jwt.MapClaims, key string) bool {
	val, ok := claims[key].(bool)
	return ok && val
}

func getIntClaim(claims jwt.MapClaims, key string) int {
	val, ok := claims[key].(float64)
	if ok {
		return int(val)
	}
	return -1
}

func getStringClaim(claims jwt.MapClaims, key string) string {
	val, ok := claims[key].(string)
	if ok {
		return val
	}
	return ""
}

func getMapClaim(claims jwt.MapClaims, key string) any {
	val, ok := claims[key]
	if ok {
		return val
	}
	return nil
}
