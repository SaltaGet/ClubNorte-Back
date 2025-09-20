package services

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

func (i *IncomeSportCourtService) IncomeSportCourtGetByID(pointSaleID, id uint) (*schemas.IncomeSportsCourtsResponse, error) {
	return nil, nil
}

func (i *IncomeSportCourtService) IncomeSportCourtGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*schemas.IncomeSportsCourtsResponseDTO, int64, error) {
	return nil, 0, nil
}

func (i *IncomeSportCourtService) IncomeSportCourtCreate(userID, pointSaleID uint, incomeCreate *schemas.IncomeSportsCourtsCreate) (uint, error) {
	return 0, nil
}

func (i *IncomeSportCourtService) IncomeSportCourtUpdate(userID, pointSaleID uint, incomeUpdate *schemas.IncomeSportsCourtsUpdate) error {
 return nil
}

func (i *IncomeSportCourtService) IncomeSportCourtDelete(pointSaleID, id uint) error {
	return nil
}