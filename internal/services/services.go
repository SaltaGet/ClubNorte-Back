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

type IncomeSportCourtService struct {
	IncomeSportCourtRepository ports.IncomeSportCourtRepository
}

type IncomeService struct {
	IncomeRepository ports.IncomeRepository
}

type MovementStockService struct {
	MovementStockRepository ports.MovementStockRepository
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

type RoleService struct {
	RoleRepository ports.RoleRepository
}

type SportCourtService struct {
	SportCourtRepository ports.SportCourtRepository
}

type StockService struct {
	StockDepositeRepository ports.StockDepositeRepository
	StockPointSaleRepository ports.StockPointSaleRepository
}

type UserService struct {
	UserRepository ports.UserRepository
}