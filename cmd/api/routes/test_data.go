package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func TestDataRoutes(app *fiber.App) {
	testData := app.Group("/api/v1/test_data")

	testData.Post("/create", controllers.TestDataCreate)
	testData.Delete("/delete", controllers.TestDataDelete)
}