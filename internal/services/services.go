package services

import "github.com/DanielChachagua/Club-Norte-Back/internal/ports"

type AuthService struct {
	AuthRepository ports.AuthRepository
	UserRepository ports.UserRepository
	PointSaleRepository ports.PointSaleRepository
}

type CategoryService struct {
	CategoryRepository ports.CategoryRepository
}

type DepositService struct {
	DepositRepository ports.DepositRepository
}

type ExpenseService struct {
	ExpenseRepository ports.ExpenseRepository
}

type ExpenseBuyService struct {
	ExpenseBuyRepository ports.ExpenseBuyRepository
}

type IncomeSportCourtService struct {
	IncomeSportCourtRepository ports.IncomeSportCourtRepository
}

type IncomeService struct {
	IncomeRepository ports.IncomeRepository
}

type MovementStockService struct {
	MovementStockRepository ports.MovementStockRepository
}

type NotificationService struct {
	NotificationRepository ports.NotificationRepository
}

type PointSaleService struct {
	PointSaleRepository ports.PointSaleRepository
}

type ProductService struct {
	ProductRepository ports.ProductRepository
}

type RegisterService struct {
	RegisterRepository ports.RegisterRepository
}

type ReportService struct {
	ReportRepository ports.ReportRepository
}

type RoleService struct {
	RoleRepository ports.RoleRepository
}

type SportCourtService struct {
	SportCourtRepository ports.SportCourtRepository
}

type StockService struct {
	StockPointSaleRepository ports.StockPointSaleRepository
}

type UserService struct {
	RoleRepository ports.RoleRepository
	UserRepository ports.UserRepository
}