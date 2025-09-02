package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type DepositRepository interface {
	DepositGetByID(id uint) (*models.Product, error)
	DepositGetByCode(code string) (*models.Product, error)
	DepositGetByName(name string) ([]*models.Product, error)
	DepositGetAll(page, limit int) ([]*models.Product, int64,error)
	DepositUpdateStock(updateStock schemas.DepositUpdateStock) (error)
}

type DepositService interface {
	DepositGetByID(id uint) (*schemas.DepositResponse, error)
	DepositGetByCode(code string) (*schemas.DepositResponse, error)
	DepositGetByName(name string) ([]*schemas.DepositResponse, error)
	DepositGetAll(page, limit int) ([]*schemas.DepositResponse, int64, error)
	DepositUpdateStock(updateStock schemas.DepositUpdateStock) (error)
}