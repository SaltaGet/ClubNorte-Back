package repositories

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

func (r *MainRepository) IncomeSportCourtGetByID(pointSaleID, id uint) (*models.IncomeSportsCourts, error) {
	return nil, nil
}

func (r *MainRepository) IncomeSportCourtGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*models.IncomeSportsCourts, int64, error) {
	return nil, 0, nil
}

func (r *MainRepository) IncomeSportCourtCreate(userID, pointSaleID uint, incomeCreate *schemas.IncomeSportsCourtsCreate) (uint, error) {
	return 0, nil
}

func (r *MainRepository) IncomeSportCourtUpdate(userID, pointSaleID uint, incomeUpdate *schemas.IncomeSportsCourtsUpdate) error {
 return nil
}

func (r *MainRepository) IncomeSportCourtDelete(pointSaleID, id uint) error {
	return nil
}
