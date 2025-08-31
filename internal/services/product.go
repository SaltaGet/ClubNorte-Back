package services

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/jinzhu/copier"
)

func (s *ProductService) ProductGetByID(id uint) (*schemas.ProductResponse, error) {
	product, err := s.ProductRepository.ProductGetByID(id)
	if err != nil {
		return nil, err
	}

	var productResponse schemas.ProductResponse
	_ = copier.Copy(&productResponse, &product)

	return &productResponse, nil
}

func (s *ProductService) ProductGetByCode(code string) (*schemas.ProductResponse, error) {
	product, err := s.ProductRepository.ProductGetByCode(code)
	if err != nil {
		return nil, err
	}

	var productResponse schemas.ProductResponse
	_ = copier.Copy(&productResponse, &product)

	return &productResponse, nil
}

func (s *ProductService) ProductGetByName(name string) ([]*schemas.ProductResponseDTO, error) {
	products, err := s.ProductRepository.ProductGetByName(name)
	if err != nil {
		return nil, err
	}

	var productsResponse []*schemas.ProductResponseDTO
	_ = copier.Copy(&productsResponse, &products)

	return productsResponse, nil
}

func (s *ProductService) ProductGetByCategoryID(categoryID uint) ([]*schemas.ProductSimpleResponse, error) {
	products, err := s.ProductRepository.ProductGetByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}

	var productsResponse []*schemas.ProductSimpleResponse
	_ = copier.Copy(&productsResponse, &products)

	return productsResponse, nil
}

func (s *ProductService) ProductGetAll(page, limit int) ([]*schemas.ProductResponseDTO, int64, error) {
	products, total, err := s.ProductRepository.ProductGetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var productsResponse []*schemas.ProductResponseDTO
	_ = copier.Copy(&productsResponse, &products)

	return productsResponse, total, nil
}

func (s *ProductService) ProductCreate(productCreate *schemas.ProductCreate) (uint, error) {
	return s.ProductRepository.ProductCreate(productCreate)
}

func (s *ProductService) ProductUpdate(productUpdate *schemas.ProductUpdate) error {
	return s.ProductRepository.ProductUpdate(productUpdate)
}

func (s *ProductService) ProductDelete(id uint) error {
	return s.ProductRepository.ProductDelete(id)
}