package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func PointSaleRoutes(app *fiber.App, controllers *controllers.PointSaleController) {
	pointSale := app.Group("/api/v1/point_sale")

	pointSale.Post("/create", controllers.PointSaleCreate, middleware.AuthMiddleware(), middleware.IsAdmin())
	pointSale.Get("/get_all", controllers.PointSaleGetAll)
	pointSale.Put("/update", controllers.PointSaleUpdate, middleware.AuthMiddleware(), middleware.IsAdmin())
	pointSale.Get("/get/:id", controllers.PointSaleGet)
	pointSale.Delete("/delete/:id", controllers.PointSaleDelete, middleware.AuthMiddleware(), middleware.IsAdmin())
}