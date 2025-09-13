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
	ProductGetAll(page, limit int) ([]*models.Product, int64, error)
	ProductCreate(productCreate *schemas.ProductCreate) (uint, error)
	ProductUpdate(productUpdate *schemas.ProductUpdate) error
	ProductDelete(id uint) error
}

type ProductService interface {
	ProductGetByID(id uint) (*schemas.ProductFullResponse, error)
	ProductGetByCode(code string) (*schemas.ProductFullResponse, error)
	ProductGetByName(name string) ([]*schemas.ProductFullResponse, error)
	ProductGetByCategoryID(categoryID uint) ([]*schemas.ProductFullResponse, error)
	ProductGetAll(page, limit int) ([]*schemas.ProductFullResponse, int64, error)
	ProductCreate(productCreate *schemas.ProductCreate) (uint, error)
	ProductUpdate(productUpdate *schemas.ProductUpdate) error
	ProductDelete(id uint) error
}