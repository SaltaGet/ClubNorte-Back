package repositories

import (
	"errors"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) RoleGetByID(id uint) (*models.Role, error) {
	var role *models.Role

	if err := r.DB.Where("id = ?", id).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "rol no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener el rol", err)
	}

	return role, nil
}

func (r *MainRepository) RoleGetAll() ([]*models.Role, error) {
	var roles []*models.Role

	if err := r.DB.Find(&roles).Error; err != nil {
		return nil, schemas.ErrorResponse(500, "error al obtener los roles", err)
	}

	return roles, nil
}