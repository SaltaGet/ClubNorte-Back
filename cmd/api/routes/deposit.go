package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func DepositRoutes(app *fiber.App, controllers *controllers.DepositController) {
	deposit := app.Group("/api/v1/deposit", middleware.AuthMiddleware())

	deposit.Get("/get_all", controllers.DepositGetAll)
	deposit.Put("/update_stock", controllers.DepositUpdateStock)
	deposit.Get("/get_by_name", controllers.DepositGetByName)
	deposit.Get("/get_by_code", controllers.DepositGetByCode)
	deposit.Get("/get/:id", controllers.DepositGetByID)
}