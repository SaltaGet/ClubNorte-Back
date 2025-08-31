package middleware

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/logging"
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

		email := getStringClaim(mapClaims, "email")
		if email == "" {
			logging.ERROR("Error: Email no proporcionado")
			return c.Status(fiber.StatusUnauthorized).JSON(schemas.Response{
				Status:  false,
				Body:    nil,
				Message: "Email no proporcionado",
			})
		}

		deps := c.Locals("deps").(*dependencies.MainContainer)
		user, err := deps.AuthController.AuthService.AuthUser(email)
		if err != nil {
			logging.ERROR("Error: %s", err.Error())
			return c.Status(fiber.StatusUnauthorized).JSON(schemas.Response{
				Status:  false,
				Body:    nil,
				Message: "Error al obtener el usuario",
			})
		}

		pointSaleID := getIntClaim(mapClaims, "point_sale_id")
		pointSale := getStringClaim(mapClaims, "point_sale")

		c.Locals("user", user)

		if pointSaleID == -1 || pointSale == "" {
			c.Locals("point_sale", nil)
		} else {
			c.Locals("point_sale", &schemas.PointSaleContext{
				ID:   2,
				Name: "mama coco",
			})
		}

		if user.IsAdmin {
			return c.Next()
		}

		if !user.IsActive {
			logging.ERROR("Error: El usuario no está activo")
			return c.Status(403).JSON(schemas.Response{
				Status:  false,
				Body:    nil,
				Message: "Usuario no activo",
			})
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
