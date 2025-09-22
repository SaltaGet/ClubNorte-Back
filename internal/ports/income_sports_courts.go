package ports

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type IncomeSportCourtRepository interface {
	IncomeSportCourtGetByID(pointSaleID, id uint) (*models.IncomeSportsCourts, error)
	IncomeSportCourtGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*models.IncomeSportsCourts, int64, error)
	IncomeSportCourtCreate(userID, pointSaleID uint, incomeCreate *schemas.IncomeSportsCourtsCreate) (uint, error)
	IncomeSportCourtUpdate(userID, pointSaleID uint, incomeUpdate *schemas.IncomeSportsCourtsUpdate) error
	IncomeSportCourtUpdatePay(userID, pointSaleID uint, incomeUpdate *schemas.IncomeSportsCourtsRestPay) error
	IncomeSportCourtDelete(pointSaleID, id uint) error
}

type IncomeSportCourtService interface {
	IncomeSportCourtGetByID(pointSaleID, id uint) (*schemas.IncomeSportsCourtsResponse, error)
	IncomeSportCourtGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*schemas.IncomeSportsCourtsResponseDTO, int64, error)
	IncomeSportCourtCreate(userID, pointSaleID uint, incomeCreate *schemas.IncomeSportsCourtsCreate) (uint, error)
	IncomeSportCourtUpdate(userID, pointSaleID uint, incomeUpdate *schemas.IncomeSportsCourtsUpdate) error
	IncomeSportCourtUpdatePay(userID, pointSaleID uint, incomeUpdate *schemas.IncomeSportsCourtsRestPay) error
	IncomeSportCourtDelete(pointSaleID, id uint) error
}
