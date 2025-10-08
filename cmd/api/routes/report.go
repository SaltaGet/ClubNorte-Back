package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func ReportRoutes(app *fiber.App, controllers *controllers.ReportController) {
	report := app.Group("/api/v1/report")
	report.Get("/get", controllers.ReportExcelGet)
	report.Post("/get_profitable_products", controllers.ReportProfitableProducts)
	report.Post("/get_by_date", controllers.ReportMovementByDate)
}