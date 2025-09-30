package services

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/jinzhu/copier"
)

func (e *ExpenseService) ExpenseGetByID(pointSaleID, id uint) (*schemas.ExpenseResponse, error) {
	expense, err := e.ExpenseRepository.ExpenseGetByID(pointSaleID, id)
	if err != nil {
		return nil, err
	}

	var expenseResponse schemas.ExpenseResponse
	_ = copier.Copy(&expenseResponse, &expense)

	var userResponse schemas.UserSimpleDTO
	_ = copier.Copy(&userResponse, &expense.User)
	expenseResponse.User = userResponse

	return &expenseResponse, nil
}

func (e *ExpenseService) ExpenseGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*schemas.ExpenseResponseDTO, int64, error) {
	expenses, total, err := e.ExpenseRepository.ExpenseGetByDate(pointSaleID, fromDate, toDate, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var expenseResponses []*schemas.ExpenseResponseDTO
	
	for _, expense := range expenses {
		var expenseResponse schemas.ExpenseResponseDTO
		_ = copier.Copy(&expenseResponse, &expense)

		expenseResponses = append(expenseResponses, &expenseResponse)
	}

	return expenseResponses, total, nil
}

func (e *ExpenseService) ExpenseCreate(userID, pointSaleID uint, expenseCreate *schemas.ExpenseCreate) (uint, error) {
	return e.ExpenseRepository.ExpenseCreate(userID, pointSaleID, expenseCreate)
}

func (e *ExpenseService) ExpenseDelete(pointSaleID, id uint) error {
	return e.ExpenseRepository.ExpenseDelete(pointSaleID, id)
}