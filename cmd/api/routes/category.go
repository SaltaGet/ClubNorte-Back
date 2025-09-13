package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App, controllers *controllers.CategoryController) {
	category := app.Group("/api/v1/category",  middleware.AuthMiddleware())

	category.Post("/create", controllers.CategoryCreate)
	category.Get("/get_all", controllers.CategoryGetAll)
	category.Put("/update", controllers.CategoryUpdate)
	category.Get("/get/:id", controllers.CategoryGet)
	category.Delete("/delete/:id", controllers.CategoryDelete)
}