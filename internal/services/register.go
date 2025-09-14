package services

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

func (r *RegisterService) RegisterOpen(pointSaleID uint, userID uint, amountOpen schemas.RegisterOpen) error {
	return nil
}

func (r *RegisterService) RegisterClose(pointSaleID uint, userID uint, amountOpen schemas.RegisterClose) error {
	return nil
}

func (r *RegisterService) RegisterInform(pointSaleID uint, userID uint, dateInform time.Time) (*schemas.RegisterInform, error) {
	return nil, nil
}