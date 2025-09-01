package services

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/jinzhu/copier"
)

func (s *DepositService) DepositGetByID(id uint) (*schemas.DepositResponse, error) {
	product, err := s.DepositRepository.DepositGetByID(id)
	if err != nil {
		return nil, err
	}

	var productResponse *schemas.DepositResponse
	_ = copier.Copy(&productResponse, &product)

	return productResponse, nil
}

func (s *DepositService) DepositGetByCode(code string) (*schemas.DepositResponse, error) {
	product, err := s.DepositRepository.DepositGetByCode(code)
	if err != nil {
		return nil, err
	}

	var productResponse *schemas.DepositResponse
	_ = copier.Copy(&productResponse, &product)

	return productResponse, nil
}

func (s *DepositService) DepositGetByName(name string) (*schemas.DepositResponse, error) {
	products, err := s.DepositRepository.DepositGetByName(name)
	if err != nil {
		return nil, err
	}

	var productsResponse []*schemas.DepositResponse
	_ = copier.Copy(&productsResponse, &products)

	return productsResponse[0], nil
}

func (s *DepositService) DepositGetAll(page, limit int) ([]*schemas.DepositResponse, int64, error) {
	products, total, err := s.DepositRepository.DepositGetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var productsResponse []*schemas.DepositResponse
	_ = copier.Copy(&productsResponse, &products)

	return productsResponse, total, nil
}

func (s *DepositService) DepositUpdateStock(productID uint, stock float64, method string) (error) {
	return s.DepositRepository.DepositUpdateStock(productID, stock, method)
}
