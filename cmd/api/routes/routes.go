package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/dependencies"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, deps *dependencies.MainContainer) {
	AuthRoutes(app, deps.AuthController)

	CategoryRoutes(app, deps.CategoryController)
	DepositRoutes(app, deps.DepositController)
	MovementStockRoutes(app, deps.MovementStockController)

	PointSaleRoutes(app, deps.PointSaleController)
	ProductRoutes(app, deps.ProductController)
	RoleRoutes(app, deps.RoleController)
	SportCourtsRoutes(app, deps.SportCourtsController)
	StockProductRoutes(app, deps.StockController)
	UserRoutes(app, deps.UserController)

	TestDataRoutes(app)
}