package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func MovementStockRoutes(app *fiber.App, controllers *controllers.MovementStockController) {
	movementStock := app.Group("/v1/movement_stock", middleware.AuthMiddleware())

	movementStock.Post("/move", controllers.MoveStock)
	movementStock.Get("/get_all", controllers.MovementStockGetAll)
	movementStock.Get("/get/:id", controllers.MovementStockGet)
}