package services

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/jinzhu/copier"
)

func (r *RegisterService) RegisterExistOpen(pointSaleID uint) (bool, error) {
	return r.RegisterRepository.RegisterExistOpen(pointSaleID)
}

func (r *RegisterService) RegisterOpen(pointSaleID uint, userID uint, amountOpen schemas.RegisterOpen) error {
	return r.RegisterRepository.RegisterOpen(pointSaleID, userID, amountOpen)
}

func (r *RegisterService) RegisterClose(pointSaleID uint, userID uint, amountOpen schemas.RegisterClose) error {
	return r.RegisterRepository.RegisterClose(pointSaleID, userID, amountOpen)
}

func (r *RegisterService) RegisterInform(pointSaleID uint, userID uint, fromDate, toDate time.Time) ([]*schemas.RegisterInformResponse, error) {
	informs, err := r.RegisterRepository.RegisterInform(pointSaleID, userID, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	
	var informResponses []*schemas.RegisterInformResponse
	_ = copier.Copy(&informResponses, &informs)

	return informResponses, nil
}