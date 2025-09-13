package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func SportCourtsRoutes(app *fiber.App, controllers *controllers.SportCourtController) {
	sportCourt := app.Group("/api/v1/sport_court", middleware.AuthMiddleware())

	sportCourt.Post("/create", controllers.SportCourtCreate)
	sportCourt.Get("/get_all", controllers.SportCourtGetAll)
	sportCourt.Put("/update", controllers.SportCourtUpdate)
	sportCourt.Get("/get_by_code", controllers.SportCourtGetByCode)
	sportCourt.Get("/get/:id", controllers.SportCourtGetByID)
	sportCourt.Delete("/delete/:id", controllers.SportCourtDelete)
}