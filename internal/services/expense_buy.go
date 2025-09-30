package services

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/jinzhu/copier"
)

func (e *ExpenseBuyService) ExpenseBuyGetByID(id uint) (*schemas.ExpenseBuyResponse, error) {
	expenseBuy, err := e.ExpenseBuyRepository.ExpenseBuyGetByID(id)
	if err != nil {
		return nil, err
	}

	var expenseBuyResponse schemas.ExpenseBuyResponse
	_ = copier.Copy(&expenseBuyResponse, &expenseBuy)

	// var userResponse schemas.UserSimpleDTO
	// _ = copier.Copy(&userResponse, &expenseBuy.User)
	// expenseBuyResponse.User = userResponse

	return &expenseBuyResponse, nil
}

func (e *ExpenseBuyService) ExpenseBuyGetByDate(fromDate, toDate time.Time, page, limit int) ([]*schemas.ExpenseBuyResponseDTO, int64, error) {
	expensesBuy, total, err := e.ExpenseBuyRepository.ExpenseBuyGetByDate(fromDate, toDate, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var expensesBuyResponses []*schemas.ExpenseBuyResponseDTO
	
	for _, expense := range expensesBuy {
		var expenseBuyResponse schemas.ExpenseBuyResponseDTO
		_ = copier.Copy(&expenseBuyResponse, &expense)

		// var userResponse schemas.UserSimpleDTO
		// _ = copier.Copy(&userResponse, &income.User)
		// incomeResponse.UserResponse = userResponse

		expensesBuyResponses = append(expensesBuyResponses, &expenseBuyResponse)
	}

	return expensesBuyResponses, total, nil
}

func (e *ExpenseBuyService) ExpenseBuyCreate(userID uint, expenseBuyCreate *schemas.ExpenseBuyCreate) (uint, error) {
	return e.ExpenseBuyRepository.ExpenseBuyCreate(userID, expenseBuyCreate)
}

func (e *ExpenseBuyService) ExpenseBuyDelete(id uint) error {
	return e.ExpenseBuyRepository.ExpenseBuyDelete(id)
}