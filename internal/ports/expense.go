package ports

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type ExpenseRepository interface {
	ExpenseGetByID(pointSaleID, id uint) (*models.Expense, error)
	ExpenseGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*models.Expense, int64, error)
	ExpenseCreate(userID, pointSaleID uint, incomeCreate *schemas.ExpenseCreate) (uint, error)
	ExpenseDelete(pointSaleID, id uint) error
}

type ExpenseService interface {
	ExpenseGetByID(pointSaleID, id uint) (*schemas.ExpenseResponse, error)
	ExpenseGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*schemas.ExpenseResponseDTO, int64, error)
	ExpenseCreate(userID, pointSaleID uint, incomeCreate *schemas.ExpenseCreate) (uint, error)
	ExpenseDelete(pointSaleID, id uint) error
}
