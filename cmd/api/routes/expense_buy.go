package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func ExpenseBuyRoutes(app *fiber.App, controllers *controllers.ExpenseBuyController) {
	expenseBuy := app.Group("/api/v1/expense_buy", middleware.AuthMiddleware())

	expenseBuy.Post("/create", controllers.ExpenseBuyCreate)
	expenseBuy.Post("/get_by_date", controllers.ExpenseBuyGetByDate)
	expenseBuy.Get("/get/:id", controllers.ExpenseBuyGetByID)
	expenseBuy.Delete("/delete/:id", controllers.ExpenseBuyDelete)
}