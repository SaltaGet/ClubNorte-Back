package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, controllers *controllers.AuthController) {
	auth := app.Group("/v1/auth")

	auth.Post("/login", controllers.Login)
}