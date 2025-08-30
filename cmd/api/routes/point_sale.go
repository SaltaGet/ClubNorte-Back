package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func PointSaleRoutes(app *fiber.App, controllers *controllers.PointSaleController) {
	pointSale := app.Group("/v1/point_sale", middleware.AuthMiddleware(), middleware.IsAdmin())

	pointSale.Post("/create", controllers.PointSaleCreate)
	pointSale.Get("/get_all", controllers.PointSaleGetAll)
	pointSale.Put("/update", controllers.PointSaleUpdate)
	pointSale.Get("/get/:id", controllers.PointSaleGet)
	pointSale.Delete("/delete/:id", controllers.PointSaleDelete)
}