package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func InformRoutes(app *fiber.App, controllers *controllers.InformController) {
	inform := app.Group("/api/v1/inform")
	inform.Get("/get", controllers.InformGet)
}