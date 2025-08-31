package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App, controllers *controllers.ProductController) {
	product := app.Group("/v1/product",  middleware.AuthMiddleware(), middleware.AuthPointSaleMiddleware())

	product.Post("/create", controllers.ProductCreate)
	product.Get("/get_all", controllers.ProductGetAll)
	product.Put("/update", controllers.ProductUpdate)
	product.Get("/get_by_code", controllers.ProductGetByCode)
	product.Get("/get_by_name", controllers.ProductGetByName)
	product.Get("/get/:id", controllers.ProductGetByID)
	product.Delete("/delete/:id", controllers.ProductDelete)
	product.Get("/get_by_category/:category_id", controllers.ProductGetByCategoryID)
}