package services

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

func (s *StockService) StockProductGetByID(pointSaleID, id uint) (*schemas.ProductResponse, error) {
	product, err := s.StockPointSaleRepository.StockProductGetByID(pointSaleID, id)
	if err != nil {
		return nil, err
	}

	productResponse := &schemas.ProductResponse{
		ID:          product.ID,
		Code:        product.Code,
		Name:        product.Name,
		Description: *product.Description,
		Category: schemas.CategoryResponse{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		},
		Price: product.Price,
		Stock: product.StockPointSales[0].Stock,
	}

	return productResponse, nil
}

func (s *StockService) StockProductGetByCode(pointSaleID uint, code string) (*schemas.ProductResponse, error) {
	product, err := s.StockPointSaleRepository.StockProductGetByCode(pointSaleID, code)
	if err != nil {
		return nil, err
	}

	productResponse := &schemas.ProductResponse{
		ID:          product.ID,
		Code:        product.Code,
		Name:        product.Name,
		Description: *product.Description,
		Category: schemas.CategoryResponse{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		},
		Price: product.Price,
		Stock: product.StockPointSales[0].Stock,
	}

	return productResponse, nil
}

func (s *StockService) StockProductGetByName(pointSaleID uint, name string) ([]*schemas.ProductResponseDTO, error) {
	products, err := s.StockPointSaleRepository.StockProductGetByName(pointSaleID, name)
	if err != nil {
		return nil, err
	}

	productsResponse := make([]*schemas.ProductResponseDTO, len(products))

	for i, prod := range products {
		productsResponse[i] = &schemas.ProductResponseDTO{
			ID:   prod.ID,
			Code: prod.Code,
			Name: prod.Name,
			Category: &schemas.CategoryResponse{
				ID:   prod.Category.ID,
				Name: prod.Category.Name,
			},
			Price: prod.Price,
			Stock: prod.StockPointSales[0].Stock,
		}
	}

	return productsResponse, nil
}

func (s *StockService) StockProductGetByCategoryID(pointSaleID, categoryID uint) ([]*schemas.ProductSimpleResponse, error) {
	products, err := s.StockPointSaleRepository.StockProductGetByCategoryID(pointSaleID, categoryID)
	if err != nil {
		return nil, err
	}

	productsResponse := make([]*schemas.ProductSimpleResponse, len(products))

	for i, prod := range products {
		productsResponse[i] = &schemas.ProductSimpleResponse{
			ID:   prod.ID,
			Code: prod.Code,
			Name: prod.Name,
			Price: prod.Price,
			Stock: prod.StockPointSales[0].Stock,
		}
	}

	return productsResponse, nil
}

func (s *StockService) StockProductGetAll(pointSaleID uint, page, limit int) ([]*schemas.ProductResponseDTO, int64, error) {
	products, total, err := s.StockPointSaleRepository.StockProductGetAll(pointSaleID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	productsResponse := make([]*schemas.ProductResponseDTO, len(products))

	for i, prod := range products {
		productsResponse[i] = &schemas.ProductResponseDTO{
			ID:   prod.ID,
			Code: prod.Code,
			Name: prod.Name,
			Category: &schemas.CategoryResponse{
				ID:   prod.Category.ID,
				Name: prod.Category.Name,
			},
			Price: prod.Price,
			Stock: prod.StockPointSales[0].Stock,
		}
	}

	return productsResponse, total, nil
}
