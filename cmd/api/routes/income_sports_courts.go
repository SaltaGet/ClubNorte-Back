package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func IncomeSportsCourtsRoutes(app *fiber.App, controllers *controllers.IncomeSportCourtController) {
	income := app.Group("/api/v1/income_sport_court", middleware.AuthMiddleware(), middleware.AuthPointSaleMiddleware())

	income.Put("/update", controllers.IncomeSportCourtUpdate)
	income.Put("/update_pay", controllers.IncomeSportCourtUpdatePay)
	income.Post("/create", controllers.IncomeSportCourtCreate)
	income.Post("/get_by_date", controllers.IncomeSportCourtGetByDate)
	income.Get("/get/:id", controllers.IncomeSportCourtGetByID)
	income.Delete("/delete/:id", controllers.IncomeSportCourtDelete)
}