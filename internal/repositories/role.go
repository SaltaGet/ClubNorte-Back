package repositories

import "github.com/DanielChachagua/Club-Norte-Back/internal/models"

func (r *MainRepository) RoleGetAll() ([]*models.Role, error) {
	var roles []*models.Role

	if err := r.DB.Where("name != ?", "admin").Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}