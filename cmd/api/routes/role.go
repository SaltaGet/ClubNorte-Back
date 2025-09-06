package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App, controllers *controllers.RoleController) {
	role := app.Group("/v1/role", middleware.AuthMiddleware())

	role.Get("/get_all", controllers.RoleGetAll)
}