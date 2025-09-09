package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type StockPointSaleRepository interface {
	StockProductGetByID(pointSaleID, id uint) (*models.Product, error)
	StockProductGetByCode(pointSaleID uint, code string) (*models.Product, error)
	StockProductGetByName(pointSaleID uint, name string) ([]*models.Product, error)
	StockProductGetByCategoryID(pointSaleID, categoryID uint) ([]*models.Product, error)
	StockProductGetAll(pointSaleID uint, page, limit int) ([]*models.Product, int64, error)
}

type StockPointSaleService interface {
	StockProductGetByID(pointSaleID, id uint) (*schemas.ProductResponse, error)
	StockProductGetByCode(pointSaleID uint, code string) (*schemas.ProductResponse, error)
	StockProductGetByName(pointSaleID uint, name string) ([]*schemas.ProductResponseDTO, error)
	StockProductGetByCategoryID(pointSaleID, categoryID uint) ([]*schemas.ProductSimpleResponse, error)
	StockProductGetAll(pointSaleID uint, page, limit int) ([]*schemas.ProductResponseDTO, int64, error)
}