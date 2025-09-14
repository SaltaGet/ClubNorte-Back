package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func SportCourtsRoutes(app *fiber.App, controllers *controllers.SportCourtController) {
	sportCourt := app.Group("/api/v1/sport_court", middleware.AuthMiddleware())

	sportCourt.Post("/create", middleware.AuthPointSaleMiddleware(), controllers.SportCourtCreate)
	sportCourt.Get("/get_all", controllers.SportCourtGetAll)
	sportCourt.Get("/get_all_by_point_sale", middleware.AuthPointSaleMiddleware(), controllers.SportCourtGetAllByPointSale)
	sportCourt.Put("/update", middleware.AuthPointSaleMiddleware(), controllers.SportCourtUpdate)
	sportCourt.Get("/get_by_code", middleware.AuthPointSaleMiddleware(), controllers.SportCourtGetByCode)
	sportCourt.Get("/get/:id", middleware.AuthPointSaleMiddleware(), controllers.SportCourtGetByID)
	sportCourt.Delete("/delete/:id", middleware.AuthPointSaleMiddleware(), controllers.SportCourtDelete)
}