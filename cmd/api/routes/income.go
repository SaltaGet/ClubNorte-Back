package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func IncomeRoutes(app *fiber.App, controllers *controllers.IncomeController) {
	income := app.Group("/api/v1/income", middleware.AuthMiddleware(), middleware.AuthPointSaleMiddleware())

	income.Post("/create", controllers.IncomeCreate)
	income.Post("/get_by_date", controllers.IncomeGetByDate)
	income.Get("/get/:id", controllers.IncomeGetByID)
	income.Delete("/delete/:id", controllers.IncomeDelete)
}