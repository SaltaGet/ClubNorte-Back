package middleware

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/server/logging"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/DanielChachagua/Club-Norte-Back/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(schemas.Response{Status: false, Body: nil, Message: "Token no proporcionado"})
		}

		claims, err := utils.VerifyToken(token)
		if err != nil {
			if errResp, ok := err.(*schemas.ErrorStruc); ok {
				logging.ERROR("Error: %s", errResp.Err.Error())
				return c.Status(errResp.StatusCode).JSON(schemas.Response{
					Status:  false,
					Body:    nil,
					Message: errResp.Message,
				})
			}
			logging.ERROR("Error: %s", err.Error())
			return c.Status(fiber.StatusUnauthorized).JSON(schemas.Response{
				Status:  false,
				Body:    nil,
				Message: "Token inválido",
			})
		}

		mapClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			logging.ERROR("Error: Claims inválidos")
			return c.Status(fiber.StatusUnauthorized).JSON(schemas.Response{
				Status:  false,
				Body:    nil,
				Message: "Claims inválidos",
			})
		}

		isAdmin := getBoolClaim(mapClaims, "is_admin")
		if isAdmin {
			return c.Next()
		}

		pointSale := getMapClaim(mapClaims, "point_sale")
		if pointSale != "" {
			c.Locals("user_id", getStringClaim(mapClaims, "id"))
			c.Locals("point_sale", pointSale)
			return c.Next()
		}

		c.Locals("user_id", getStringClaim(mapClaims, "id"))
		c.Locals("point_sale", nil)
		return c.Next()
	}
}

func getBoolClaim(claims jwt.MapClaims, key string) bool {
	val, ok := claims[key].(bool)
	return ok && val
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