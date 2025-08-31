package repositories

import (
	"errors"
	"fmt"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"gorm.io/gorm"
)

func (r *MainRepository) UserGetByID(id uint) (*models.User, error) {
	var user *models.User

	if err := r.DB.Preload("Role").Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("punto de venta no encontrado")
		}
		return nil, err
	}

	return user, nil
}

func (r *MainRepository) UserGetByEmail(email string) (*models.User, error) {
	var user *models.User

	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("punto de venta no encontrado")
		}
		return nil, err
	}

	return user, nil
}