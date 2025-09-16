package ports

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type IncomeRepository interface {
	IncomeGetByID(pointSaleID, id uint) (*models.Income, error)
	IncomeGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*models.Income, int64, error)
	IncomeCreate(userID, pointSaleID uint, incomeCreate *schemas.IncomeCreate) (uint, error)
	IncomeUpdate(userID, pointSaleID uint, incomeUpdate *schemas.IncomeUpdate) (error)
	IncomeDelete(pointSaleID, id uint) error
}

type IncomeService interface {
	IncomeGetByID(pointSaleID, id uint) (*schemas.IncomeResponse, error)
	IncomeGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*schemas.IncomeResponseDTO, int64, error)
	IncomeCreate(userID, pointSaleID uint, incomeCreate *schemas.IncomeCreate) (uint, error)
	IncomeUpdate(userID, pointSaleID uint, incomeUpdate *schemas.IncomeUpdate) error
	IncomeDelete(pointSaleID, id uint) error
}
