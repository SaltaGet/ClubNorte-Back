package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type PointSaleRepository interface {
	PointSaleGet(id uint) (*models.PointSale, error)
	PointSaleGetAll() ([]*models.PointSale, error)
	PointSaleCreate(pointSaleCreate *schemas.PointSaleCreate) (uint, error)
	PointSaleUpdate(pointSaleUpdate *schemas.PointSaleUpdate) error
	PointSaleDelete(id uint) error
}

type PointSaleService interface {
	PointSaleGet(id uint) (*schemas.PointSaleResponse, error)
	PointSaleGetAll() ([]*schemas.PointSaleResponse, error)
	PointSaleCreate(pointSaleCreate *schemas.PointSaleCreate) (uint, error)
	PointSaleUpdate(pointSaleUpdate *schemas.PointSaleUpdate) error
	PointSaleDelete(id uint) error
}