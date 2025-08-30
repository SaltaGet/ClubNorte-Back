package services

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/jinzhu/copier"
)

func (s *PointSaleService) PointSaleGet(id uint) (*schemas.PointSaleResponse, error) {
	pointSale, err := s.PointSaleRepository.PointSaleGet(id)
	if err != nil {
		return nil, err
	}

	var pointSaleResponse schemas.PointSaleResponse
	_ = copier.Copy(&pointSaleResponse, &pointSale)

	return &pointSaleResponse, nil
}

func (s *PointSaleService) PointSaleGetAll() ([]*schemas.PointSaleResponse, error) {
	pointSales, err := s.PointSaleRepository.PointSaleGetAll()
	if err != nil {
		return nil, err
	}

	var pointSalesResponse []*schemas.PointSaleResponse
	_ = copier.Copy(&pointSalesResponse, &pointSales)

	return pointSalesResponse, nil
}

func (s *PointSaleService) PointSaleCreate(pointSaleCreate *schemas.PointSaleCreate) (uint, error) {
	return s.PointSaleRepository.PointSaleCreate(pointSaleCreate)
}

func (s *PointSaleService) PointSaleUpdate(pointSaleUpdate *schemas.PointSaleUpdate) error {
	return s.PointSaleRepository.PointSaleUpdate(pointSaleUpdate)
}

func (s *PointSaleService) PointSaleDelete(id uint) error {
	return s.PointSaleRepository.PointSaleDelete(id)
}