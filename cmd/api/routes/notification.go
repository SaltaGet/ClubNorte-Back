package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	// "github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func NotificationRoutes(app *fiber.App, controllers *controllers.NotificationController) {
	notification := app.Group("/api/v1/notification") //, middleware.AuthMiddleware())

	notification.Get("/alert", controllers.NotificationAlert)
}