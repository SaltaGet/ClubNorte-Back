package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func StockProductRoutes(app *fiber.App, controllers *controllers.StockController) {
	product := app.Group("/api/v1/point_sale_product",  middleware.AuthMiddleware(), middleware.AuthPointSaleMiddleware())

	product.Get("/get_all", controllers.StockProductGetAll)
	product.Get("/get_by_code", controllers.StockProductGetByCode)
	product.Get("/get_by_name", controllers.StockProductGetByName)
	product.Get("/get/:id", controllers.StockProductGetByID)
	product.Get("/get_by_category/:category_id", controllers.StockProductGetByCategoryID)
}