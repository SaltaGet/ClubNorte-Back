package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func IncomeSportsCourtsRoutes(app *fiber.App, controllers *controllers.IncomeSportCourtController) {
	incomeSportCourt := app.Group("/api/v1/income_sport_court", middleware.AuthMiddleware(), middleware.AuthPointSaleMiddleware())

	incomeSportCourt.Post("/create", controllers.IncomeSportCourtCreate)
	incomeSportCourt.Put("/update_pay", controllers.IncomeSportCourtUpdatePay)
	incomeSportCourt.Post("/get_by_date", controllers.IncomeSportCourtGetByDate)
	incomeSportCourt.Get("/get/:id", controllers.IncomeSportCourtGetByID)
	incomeSportCourt.Delete("/delete/:id", controllers.IncomeSportCourtDelete)
}