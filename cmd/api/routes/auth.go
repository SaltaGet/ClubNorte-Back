package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, controllers *controllers.AuthController) {
	auth := app.Group("/v1/auth")

	auth.Get("/current_user", middleware.AuthMiddleware(), controllers.CurrentUser)
	auth.Get("/current_point_sale", middleware.AuthMiddleware(), middleware.AuthPointSaleMiddleware(), controllers.CurrentPointSale)
	auth.Post("/login", controllers.Login)
	auth.Post("/logout", middleware.AuthMiddleware(), controllers.Logout)
	auth.Post("/logout_point_sale",  middleware.AuthMiddleware(), middleware.AuthPointSaleMiddleware(), controllers.LogoutPointSale)
	auth.Post("/login_point_sale/:point_sale_id",  middleware.AuthMiddleware(), controllers.LoginPointSale)
}