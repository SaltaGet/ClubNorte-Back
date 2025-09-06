package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type RoleRepository interface {
	RoleGetAll() ([]*models.Role, error)
}

type RoleService interface {
	RoleGetAll() ([]*schemas.RoleResponse, error)
}