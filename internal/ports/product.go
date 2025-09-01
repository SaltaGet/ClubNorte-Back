package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type ProductRepository interface {
	ProductGetByID(id uint) (*models.Product, error)
	ProductGetByCode(code string) (*models.Product, error)
	ProductGetByName(name string) ([]*models.Product, error)
	ProductGetByCategoryID(categoryID uint) ([]*models.Product, error)
	ProductGetAll(pointSaleID uint, page, limit int) ([]*models.Product, int64, error)
	ProductCreate(productCreate *schemas.ProductCreate) (uint, error)
	ProductUpdate(productUpdate *schemas.ProductUpdate) error
	ProductDelete(id uint) error
}

type ProductService interface {
	ProductGetByID(id uint) (*schemas.ProductResponse, error)
	ProductGetByCode(code string) (*schemas.ProductResponse, error)
	ProductGetByName(name string) ([]*schemas.ProductResponseDTO, error)
	ProductGetByCategoryID(categoryID uint) ([]*schemas.ProductSimpleResponse, error)
	ProductGetAll(pointSaleID uint, page, limit int) ([]*schemas.ProductResponseDTO, int64, error)
	ProductCreate(productCreate *schemas.ProductCreate) (uint, error)
	ProductUpdate(productUpdate *schemas.ProductUpdate) error
	ProductDelete(id uint) error
}