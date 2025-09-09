package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type RoleRepository interface {
	RoleGetByID(id uint) (*models.Role, error)
	RoleGetAll() ([]*models.Role, error)
}

type RoleService interface {
	RoleGetAll() ([]*schemas.RoleResponse, error)
}