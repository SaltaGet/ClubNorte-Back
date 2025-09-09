package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, controllers *controllers.UserController) {
	user := app.Group("/v1/user", middleware.AuthMiddleware())

	user.Get("/get_all", middleware.IsAdmin(), controllers.UserGetAll)
	user.Get("/get_by_email", middleware.IsAdmin(), controllers.UserGetByEmail)
	user.Post("/create", middleware.IsAdmin(), controllers.UserCreate)
	user.Put("/update", controllers.UserUpdate)
	user.Put("/update_password", controllers.UserUpdatePassword)
	user.Delete("/delete/:id", middleware.IsAdmin(), controllers.UserDelete)
	user.Get("/get/:id", middleware.IsAdmin(), controllers.UserGetByID)
}