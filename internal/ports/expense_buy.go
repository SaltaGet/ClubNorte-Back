package ports

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type ExpenseBuyRepository interface {
	ExpenseBuyGetByID(id uint) (*models.ExpenseBuy, error)
	ExpenseBuyGetByDate(fromDate, toDate time.Time, page, limit int) ([]*models.ExpenseBuy, int64, error)
	ExpenseBuyCreate(userID uint, incomeCreate *schemas.ExpenseBuyCreate) (uint, error)
	ExpenseBuyDelete(id uint) error
}

type ExpenseBuyService interface {
	ExpenseBuyGetByID(id uint) (*schemas.ExpenseBuyResponse, error)
	ExpenseBuyGetByDate(fromDate, toDate time.Time, page, limit int) ([]*schemas.ExpenseBuyResponseDTO, int64, error)
	ExpenseBuyCreate(userID uint, incomeCreate *schemas.ExpenseBuyCreate) (uint, error)
	ExpenseBuyDelete(id uint) error
}