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

func (r *MainRepository) AuthUser(email string) (*models.User, error) {
	var user models.User

	err := r.DB.Preload("Role").
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MainRepository) AuthPointSale(userID uint, pointSaleID uint) (*models.PointSale, error) {
	var pointSale models.PointSale

	err := r.DB.Joins("JOIN user_point_sales ups ON ups.point_sale_id = point_sales.id").
    Where("ups.user_id = ?", userID).
    Find(&pointSale).Error
	if err != nil {
		return nil, err
	}

	return &pointSale, nil
}

func (r *MainRepository) LoginPointSale(userID uint, pointSaleID uint) (*models.PointSale, error) {
	var pointSale models.PointSale

	err := r.DB.Joins("JOIN user_point_sales ups ON ups.point_sale_id = point_sales.id").
    Where("ups.user_id = ?", userID).
    Find(&pointSale).Error
	if err != nil {
		return nil, err
	}

	return &pointSale, nil
}
