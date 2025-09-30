package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func ExpenseRoutes(app *fiber.App, controllers *controllers.ExpenseController) {
	expense := app.Group("/api/v1/expense", middleware.AuthMiddleware(), middleware.AuthPointSaleMiddleware())

	expense.Post("/create", controllers.ExpenseCreate)
	expense.Post("/get_by_date", controllers.ExpenseGetByDate)
	expense.Get("/get/:id", controllers.ExpenseGetByID)
	expense.Delete("/delete/:id", controllers.ExpenseDelete)
}