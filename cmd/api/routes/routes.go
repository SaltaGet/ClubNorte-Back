package routes

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/dependencies"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, deps *dependencies.MainContainer) {
	AuthRoutes(app, deps.AuthController)

	CategoryRoutes(app, deps.CategoryController)
	DepositRoutes(app, deps.DepositController)
	ExpenseRoutes(app, deps.ExpenseController)
	ExpenseBuyRoutes(app, deps.ExpenseBuyController)
	IncomeRoutes(app, deps.IncomeController)
	IncomeSportsCourtsRoutes(app, deps.IncomeSportCourtController)
	InformRoutes(app, deps.InformController)
	MovementStockRoutes(app, deps.MovementStockController)
	NotificationRoutes(app, deps.NotificationController)

	PointSaleRoutes(app, deps.PointSaleController)
	ProductRoutes(app, deps.ProductController)
	RegisterRoutes(app, deps.RegisterController)
	RoleRoutes(app, deps.RoleController)
	SportCourtsRoutes(app, deps.SportCourtsController)
	StockProductRoutes(app, deps.StockController)
	UserRoutes(app, deps.UserController)

	TestDataRoutes(app)
}