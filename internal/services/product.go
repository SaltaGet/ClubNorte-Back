package services

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

func (s *ProductService) ProductGetByID(id uint) (*schemas.ProductResponse, error) {
	product, err := s.ProductRepository.ProductGetByID(id)
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
		Stock: product.StockPointSale.Stock,
	}

	return productResponse, nil
}

func (s *ProductService) ProductGetByCode(code string) (*schemas.ProductResponse, error) {
	product, err := s.ProductRepository.ProductGetByCode(code)
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
		Stock: product.StockPointSale.Stock,
	}

	return productResponse, nil
}

func (s *ProductService) ProductGetByName(name string) ([]*schemas.ProductResponseDTO, error) {
	products, err := s.ProductRepository.ProductGetByName(name)
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
			Stock: prod.StockPointSale.Stock,
		}
	}

	return productsResponse, nil
}

func (s *ProductService) ProductGetByCategoryID(categoryID uint) ([]*schemas.ProductSimpleResponse, error) {
	products, err := s.ProductRepository.ProductGetByCategoryID(categoryID)
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
			Stock: prod.StockPointSale.Stock,
		}
	}

	return productsResponse, nil
}

func (s *ProductService) ProductGetAll(pointSaleID uint, page, limit int) ([]*schemas.ProductResponseDTO, int64, error) {
	products, total, err := s.ProductRepository.ProductGetAll(pointSaleID, page, limit)
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
			Stock: prod.StockPointSale.Stock,
		}
	}

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
