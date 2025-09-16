package services

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

func (s *ProductService) ProductGetByID(id uint) (*schemas.ProductFullResponse, error) {
	product, err := s.ProductRepository.ProductGetByID(id)
	if err != nil {
		return nil, err
	}

	var productResponse schemas.ProductFullResponse

	productResponse.ID = product.ID
	productResponse.Code = product.Code
	productResponse.Name = product.Name
	productResponse.Description = *product.Description
	productResponse.Category = schemas.CategoryResponse{
		ID:   product.Category.ID,
		Name: product.Category.Name,
	}
	productResponse.Price = product.Price
	
	if product.StockDeposit != nil {
		productResponse.StockDeposit = &schemas.StockDepositResponse{
			ID:    product.StockDeposit.ID,
			Stock: product.StockDeposit.Stock,
		}
	} else {
		productResponse.StockDeposit = &schemas.StockDepositResponse{
			ID:    0,
			Stock: 0,
		}
	}
	
	for _, stock := range product.StockPointSales {
		productResponse.StockPointSales = append(productResponse.StockPointSales, &schemas.PointSaleStock{
			ID:    stock.PointSale.ID,
			Name:  stock.PointSale.Name,
			Stock: stock.Stock,
		})
	}

	return &productResponse, nil
}

func (s *ProductService) ProductGetByCode(code string) (*schemas.ProductFullResponse, error) {
	product, err := s.ProductRepository.ProductGetByCode(code)
	if err != nil {
		return nil, err
	}

	var productResponse schemas.ProductFullResponse

	productResponse.ID = product.ID
	productResponse.Code = product.Code
	productResponse.Name = product.Name
	productResponse.Description = *product.Description
	productResponse.Category = schemas.CategoryResponse{
		ID:   product.Category.ID,
		Name: product.Category.Name,
	}
	productResponse.Price = product.Price

	if product.StockDeposit != nil {
		productResponse.StockDeposit = &schemas.StockDepositResponse{
			ID:    product.StockDeposit.ID,
			Stock: product.StockDeposit.Stock,
		}
	} else {
		productResponse.StockDeposit = &schemas.StockDepositResponse{
			ID:    0,
			Stock: 0,
		}
	}

	for _, stock := range product.StockPointSales {
		productResponse.StockPointSales = append(productResponse.StockPointSales, &schemas.PointSaleStock{
			ID:    stock.PointSale.ID,
			Name:  stock.PointSale.Name,
			Stock: stock.Stock,
		})
	}

	return &productResponse, nil
}

func (s *ProductService) ProductGetByName(name string) ([]*schemas.ProductFullResponse, error) {
	products, err := s.ProductRepository.ProductGetByName(name)
	if err != nil {
		return nil, err
	}

	productsResponse := make([]*schemas.ProductFullResponse, len(products))

	for i, prod := range products {
		productsResponse[i] = &schemas.ProductFullResponse{
			ID:   prod.ID,
			Code: prod.Code,
			Name: prod.Name,
			Category: schemas.CategoryResponse{
				ID:   prod.Category.ID,
				Name: prod.Category.Name,
			},
			Price: prod.Price,
		}
		if prod.StockDeposit != nil {
			productsResponse[i].StockDeposit = &schemas.StockDepositResponse{
				ID:    prod.StockDeposit.ID,
				Stock: prod.StockDeposit.Stock,
			}
		} else {
			productsResponse[i].StockDeposit = &schemas.StockDepositResponse{
				ID:    0,
				Stock: 0,
			}
		}
		for _, stock := range prod.StockPointSales {
			productsResponse[i].StockPointSales = append(productsResponse[i].StockPointSales, &schemas.PointSaleStock{
				ID:    stock.PointSale.ID,
				Name:  stock.PointSale.Name,
				Stock: stock.Stock,
			})
		}
	}

	return productsResponse, nil
}

func (s *ProductService) ProductGetByCategoryID(categoryID uint) ([]*schemas.ProductFullResponse, error) {
	products, err := s.ProductRepository.ProductGetByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}

	productsResponse := make([]*schemas.ProductFullResponse, len(products))

	for i, prod := range products {
		productsResponse[i] = &schemas.ProductFullResponse{
			ID:   prod.ID,
			Code: prod.Code,
			Name: prod.Name,
			Category: schemas.CategoryResponse{
				ID:   prod.Category.ID,
				Name: prod.Category.Name,
			},
			Price: prod.Price,
		}
		if prod.StockDeposit != nil {
			productsResponse[i].StockDeposit = &schemas.StockDepositResponse{
				ID:    prod.StockDeposit.ID,
				Stock: prod.StockDeposit.Stock,
			}
		} else {
			productsResponse[i].StockDeposit = &schemas.StockDepositResponse{
				ID:    0,
				Stock: 0,
			}
		}
		for _, stock := range prod.StockPointSales {
			productsResponse[i].StockPointSales = append(productsResponse[i].StockPointSales, &schemas.PointSaleStock{
				ID:    stock.PointSale.ID,
				Name:  stock.PointSale.Name,
				Stock: stock.Stock,
			})
		}
	}

	return productsResponse, nil
}

func (s *ProductService) ProductGetAll(page, limit int) ([]*schemas.ProductFullResponse, int64, error) {
	products, total, err := s.ProductRepository.ProductGetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	productsResponse := make([]*schemas.ProductFullResponse, len(products))

	for i, prod := range products {
		productsResponse[i] = &schemas.ProductFullResponse{
			ID:   prod.ID,
			Code: prod.Code,
			Name: prod.Name,
			Category: schemas.CategoryResponse{
				ID:   prod.Category.ID,
				Name: prod.Category.Name,
			},
			Price: prod.Price,
		}
		if prod.StockDeposit != nil {
			productsResponse[i].StockDeposit = &schemas.StockDepositResponse{
				ID:    prod.StockDeposit.ID,
				Stock: prod.StockDeposit.Stock,
			}
		} else {
			productsResponse[i].StockDeposit = &schemas.StockDepositResponse{
				ID:    0,
				Stock: 0,
			}
		}

		for _, stock := range prod.StockPointSales {
			productsResponse[i].StockPointSales = append(productsResponse[i].StockPointSales, &schemas.PointSaleStock{
				ID:    stock.PointSale.ID,
				Name:  stock.PointSale.Name,
				Stock: stock.Stock,
			})
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
