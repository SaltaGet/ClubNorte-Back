package services

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/jinzhu/copier"
)

func (i *IncomeSportCourtService) IncomeSportCourtGetByID(pointSaleID, id uint) (*schemas.IncomeSportsCourtsResponse, error) {
	income, err := i.IncomeSportCourtRepository.IncomeSportCourtGetByID(pointSaleID, id)
	if err != nil {
		return nil, err
	}

	var incomeResponse schemas.IncomeSportsCourtsResponse
	_ = copier.Copy(&incomeResponse, &income)

	var userResponse schemas.UserSimpleDTO
	_ = copier.Copy(&userResponse, &income.User)
	incomeResponse.UserResponse = userResponse

	return &incomeResponse, nil
}

func (i *IncomeSportCourtService) IncomeSportCourtGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*schemas.IncomeSportsCourtsResponseDTO, int64, error) {
	incomes, total, err := i.IncomeSportCourtRepository.IncomeSportCourtGetByDate(pointSaleID, fromDate, toDate, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var incomeResponses []*schemas.IncomeSportsCourtsResponseDTO
	
	for _, income := range incomes {
		var incomeResponse schemas.IncomeSportsCourtsResponseDTO
		_ = copier.Copy(&incomeResponse, &income)

		incomeResponses = append(incomeResponses, &incomeResponse)
	}

	return incomeResponses, total, nil
}

func (i *IncomeSportCourtService) IncomeSportCourtCreate(userID, pointSaleID uint, incomeCreate *schemas.IncomeSportsCourtsCreate) (uint, error) {
	return i.IncomeSportCourtRepository.IncomeSportCourtCreate(userID, pointSaleID, incomeCreate)
}

func (i *IncomeSportCourtService) IncomeSportCourtUpdatePay(userID, pointSaleID uint, incomeUpdate *schemas.IncomeSportsCourtsRestPay) error {
 return i.IncomeSportCourtRepository.IncomeSportCourtUpdatePay(userID, pointSaleID, incomeUpdate)
}

func (i *IncomeSportCourtService) IncomeSportCourtDelete(pointSaleID, id uint) error {
	return nil
}