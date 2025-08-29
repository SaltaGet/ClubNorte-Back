package repositories

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

func (r *MainRepository) Login(params *schemas.Login) (*models.User, error) {
	var user models.User

	err := r.DB.Preload("Role").
		Where("email = ?", params.Email).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
