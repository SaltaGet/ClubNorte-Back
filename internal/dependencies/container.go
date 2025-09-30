package dependencies

import (
	"github.com/DanielChachagua/Club-Norte-Back/cmd/api/controllers"
	"github.com/DanielChachagua/Club-Norte-Back/internal/repositories"
	"github.com/DanielChachagua/Club-Norte-Back/internal/services"
	"gorm.io/gorm"
)

type MainContainer struct {
	AuthController *controllers.AuthController
	CategoryController *controllers.CategoryController
	DepositController *controllers.DepositController
	ExpenseController *controllers.ExpenseController
	ExpenseBuyController *controllers.ExpenseBuyController
	IncomeSportCourtController *controllers.IncomeSportCourtController
	IncomeController *controllers.IncomeController
	InformController *controllers.InformController
	MovementStockController *controllers.MovementStockController
	NotificationController *controllers.NotificationController
	PointSaleController *controllers.PointSaleController
	ProductController *controllers.ProductController
	RegisterController *controllers.RegisterController
	RoleController *controllers.RoleController
	SportCourtsController *controllers.SportCourtController
	StockController *controllers.StockController
	UserController *controllers.UserController

	NotificationChannel chan struct{}
}

func NewMainContainer(db *gorm.DB) *MainContainer {
	repo := &repositories.MainRepository{DB: db}

	authSvc := &services.AuthService{AuthRepository: repo, UserRepository: repo, PointSaleRepository: repo}
	categorySvc := &services.CategoryService{CategoryRepository: repo}
	depositSvc := &services.DepositService{DepositRepository: repo}
	expenseSvc := &services.ExpenseService{ExpenseRepository: repo}
	expenseBuySvc := &services.ExpenseBuyService{ExpenseBuyRepository: repo}
	incomeSportCourtSvc := &services.IncomeSportCourtService{IncomeSportCourtRepository: repo}
	incomeSvc := &services.IncomeService{IncomeRepository: repo}
	informSvc := &services.InformService{InformRepository: repo}
	movementStockSvc := &services.MovementStockService{MovementStockRepository: repo}
	notificationSvc := &services.NotificationService{NotificationRepository: repo}
	pointSaleSvc := &services.PointSaleService{PointSaleRepository: repo}
	productSvc := &services.ProductService{ProductRepository: repo}
	registerSvc := &services.RegisterService{RegisterRepository: repo}
	roleSvc := &services.RoleService{RoleRepository: repo}
	sportCourtSvc := &services.SportCourtService{SportCourtRepository: repo}
	stockSvc := &services.StockService{StockPointSaleRepository: repo}
	userSvc := &services.UserService{UserRepository: repo, RoleRepository: repo}

	notificationCh := make(chan struct{}, 100) // Buffer más grande para múltiples notificaciones

	// Crear NotificationController primero
	notificationCtrl := &controllers.NotificationController{
		NotificationService: notificationSvc,
		NotifyCh:           notificationCh,
	}

	movementStockCtrl := &controllers.MovementStockController{
		MovementStockService:   movementStockSvc,
		NotificationController: notificationCtrl,
	}

	return &MainContainer{
		AuthController: &controllers.AuthController{AuthService: authSvc},
		CategoryController: &controllers.CategoryController{CategoryService: categorySvc},
		DepositController: &controllers.DepositController{DepositService: depositSvc},
		ExpenseController: &controllers.ExpenseController{ExpenseService: expenseSvc},
		ExpenseBuyController: &controllers.ExpenseBuyController{ExpenseBuyService: expenseBuySvc},
		IncomeSportCourtController: &controllers.IncomeSportCourtController{IncomeSportCourtService: incomeSportCourtSvc},
		IncomeController: &controllers.IncomeController{IncomeService: incomeSvc},
		InformController: &controllers.InformController{InformService: informSvc},
		MovementStockController: movementStockCtrl,
		NotificationController: notificationCtrl,
		PointSaleController: &controllers.PointSaleController{PointSaleService: pointSaleSvc},
		ProductController: &controllers.ProductController{ProductService: productSvc},
		RegisterController: &controllers.RegisterController{RegisterService: registerSvc},
		RoleController: &controllers.RoleController{RoleService: roleSvc},
		SportCourtsController: &controllers.SportCourtController{SportCourtService: sportCourtSvc},
		StockController: &controllers.StockController{StockService: stockSvc},
		UserController: &controllers.UserController{UserService: userSvc},

		NotificationChannel: notificationCh,
	}
}
