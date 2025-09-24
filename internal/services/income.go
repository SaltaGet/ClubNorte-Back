package services

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/jinzhu/copier"
)

func (i *IncomeService) IncomeGetByID(pointSaleID, id uint) (*schemas.IncomeResponse, error) {
	income, err := i.IncomeRepository.IncomeGetByID(pointSaleID, id)
	if err != nil {
		return nil, err
	}

	var incomeResponse schemas.IncomeResponse
	_ = copier.Copy(&incomeResponse, &income)

	var userResponse schemas.UserSimpleDTO
	_ = copier.Copy(&userResponse, &income.User)
	incomeResponse.UserResponse = userResponse

	return &incomeResponse, nil
}

func (i *IncomeService) IncomeGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*schemas.IncomeResponseDTO, int64, error) {
	incomes, total, err := i.IncomeRepository.IncomeGetByDate(pointSaleID, fromDate, toDate, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var incomeResponses []*schemas.IncomeResponseDTO
	
	for _, income := range incomes {
		var incomeResponse schemas.IncomeResponseDTO
		_ = copier.Copy(&incomeResponse, &income)

		var userResponse schemas.UserSimpleDTO
		_ = copier.Copy(&userResponse, &income.User)
		incomeResponse.UserResponse = userResponse

		incomeResponses = append(incomeResponses, &incomeResponse)
	}

	return incomeResponses, total, nil
}

func (i *IncomeService) IncomeCreate(userID, pointSaleID uint, incomeCreate *schemas.IncomeCreate) (uint, error) {
	return i.IncomeRepository.IncomeCreate(userID, pointSaleID, incomeCreate)
}

func (i *IncomeService) IncomeDelete(pointSaleID, id uint) error {
	return i.IncomeRepository.IncomeDelete(pointSaleID, id)
}