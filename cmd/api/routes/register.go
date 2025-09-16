package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, controllers *controllers.RegisterController) {
	register := app.Group("/api/v1/register", middleware.AuthMiddleware(), middleware.AuthPointSaleMiddleware())

	register.Get("/exist_open", controllers.RegisterExistOpen)
	register.Post("/open", controllers.RegisterOpen)
	register.Post("/inform", controllers.RegiterInform)
	register.Post("/close", controllers.RegisterClose)
}