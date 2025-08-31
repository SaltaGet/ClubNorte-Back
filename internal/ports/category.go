package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)


type CategoryRepository interface {
	CategoryGetByID(id uint) (*models.Category, error)
	CategoryGetAll() ([]*models.Category, error)
	CategoryCreate(categoryCreate *schemas.CategoryCreate) (uint, error)
	CategoryUpdate(categoryUpdate *schemas.CategoryUpdate) error
	CategoryDelete(id uint) error
}

type CategoryService interface {
	CategoryGetByID(id uint) (*schemas.CategoryResponse, error)
	CategoryGetAll() ([]*schemas.CategoryResponse, error)
	CategoryCreate(categoryCreate *schemas.CategoryCreate) (uint, error)
	CategoryUpdate(categoryUpdate *schemas.CategoryUpdate) error
	CategoryDelete(id uint) error
}