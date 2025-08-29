package dependencies

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/server/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/internal/repositories"
	"github.com/DanielChachagua/Club-Norte-Back/internal/services"
	"gorm.io/gorm"
)

type MainContainer struct {
	AuthController *controllers.AuthController
	CategoryController *controllers.CategoryController
	DepositController *controllers.DepositController
	ExpenseController *controllers.ExpenseController
	IncomeSportCourtController *controllers.IncomeSportCourtController
	IncomeController *controllers.IncomeController
	MovementStockController *controllers.MovementStockController
	PointSaleController *controllers.PointSaleController
	ProductController *controllers.ProductController
	RegisterController *controllers.RegisterController
	RoleController *controllers.RoleController
	SportCourtsController *controllers.SportCourtController
	// StockDepositeController *controllers.StockDepositeController
	StockController *controllers.StockController
	UserController *controllers.UserController
}

func NewMainContainer(db *gorm.DB) *MainContainer {
	repo := &repositories.MainRepository{DB: db}

	authSvc := &services.AuthService{AuthRepository: repo}
	categorySvc := &services.CategoryService{CategoryRepository: repo}
	depositSvc := &services.DepositService{DepositRepository: repo}
	expenseSvc := &services.ExpenseService{ExpenseRepository: repo}
	incomeSportCourtSvc := &services.IncomeSportCourtService{IncomeSportCourtRepository: repo}
	incomeSvc := &services.IncomeService{IncomeRepository: repo}
	movementStockSvc := &services.MovementStockService{MovementStockRepository: repo}
	pointSaleSvc := &services.PointSaleService{PointSaleRepository: repo}
	productSvc := &services.ProductService{ProductRepository: repo}
	registerSvc := &services.RegisterService{RegisterRepository: repo}
	roleSvc := &services.RoleService{RoleRepository: repo}
	sportCourtSvc := &services.SportCourtService{SportCourtRepository: repo}
	stockSvc := &services.StockService{StockDepositeRepository: repo, StockPointSaleRepository: repo}
	// stockPointSaleSvc := &services.CategoryService{StockPointSaleRepository: repo}
	userSvc := &services.UserService{UserRepository: repo}

	return &MainContainer{
		AuthController: &controllers.AuthController{AuthService: authSvc},
		CategoryController: &controllers.CategoryController{CategoryService: categorySvc},
		DepositController: &controllers.DepositController{DepositService: depositSvc},
		ExpenseController: &controllers.ExpenseController{ExpenseService: expenseSvc},
		IncomeSportCourtController: &controllers.IncomeSportCourtController{IncomeSportCourtService: incomeSportCourtSvc},
		IncomeController: &controllers.IncomeController{IncomeService: incomeSvc},
		MovementStockController: &controllers.MovementStockController{MovementStockService: movementStockSvc},
		PointSaleController: &controllers.PointSaleController{PointSaleService: pointSaleSvc},
		ProductController: &controllers.ProductController{ProductService: productSvc},
		RegisterController: &controllers.RegisterController{RegisterService: registerSvc},
		RoleController: &controllers.RoleController{RoleService: roleSvc},
		SportCourtsController: &controllers.SportCourtController{SportCourtService: sportCourtSvc},
		StockController: &controllers.StockController{StockService: stockSvc},
		// StockPointSaleController: &controllers.StockPointSaleController{StockPointSaleService: stockPointSaleSvc},
		UserController: &controllers.UserController{UserService: userSvc},
	}
}
