package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, controllers *controllers.UserController) {
	user := app.Group("/v1/user", middleware.AuthMiddleware())

	user.Get("/get_all", controllers.UserGetAll)
	user.Post("/create", controllers.UserCreate)
	user.Put("/update", controllers.UserUpdate)
	user.Put("/update_password", controllers.UserUpdatePassword)
	user.Delete("/delete/:id", controllers.UserDelete)
	user.Get("/get/:id", controllers.UserGetByID)
}