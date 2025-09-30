package controllers

import "github.com/DanielChachagua/Club-Norte-Back/internal/ports"

type AuthController struct {	
	AuthService ports.AuthService
}

type CategoryController struct {	
	CategoryService ports.CategoryService
}

type DepositController struct {
	DepositService ports.DepositService
}

type ExpenseController struct {
	ExpenseService ports.ExpenseService
}

type ExpenseBuyController struct {
	ExpenseBuyService ports.ExpenseBuyService
}

type IncomeSportCourtController struct {
	IncomeSportCourtService ports.IncomeSportCourtService
}

type IncomeController struct {
	IncomeService ports.IncomeService
}

type InformController struct {
	InformService ports.InformService
}

type MovementStockController struct {
	MovementStockService ports.MovementStockService
	NotificationController *NotificationController
}

type NotificationController struct {
	NotificationService ports.NotificationService
	NotifyCh         chan struct{}
}

type PointSaleController struct {
	PointSaleService ports.PointSaleService
}

type ProductController struct {
	ProductService ports.ProductService
}

type RegisterController struct {
	RegisterService ports.RegisterService
}

type RoleController struct {
	RoleService ports.RoleService
}

type SportCourtController struct {
	SportCourtService ports.SportCourtService
}

type StockController struct {
	StockService ports.StockPointSaleService
}

// type StockPointSaleController struct {
// 	StockPointSaleService ports.StockPointSaleService
// }

type UserController struct {
	UserService ports.UserService
}