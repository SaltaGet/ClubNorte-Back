package repositories

import (
	"errors"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) Login(params *schemas.Login) (*models.User, error) {
	var user models.User

	err := r.DB.Preload("Role").
		Where("email = ?", params.Email).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(401, "credenciales incorrectas", err)
		}
		return nil, schemas.ErrorResponse(500, "error al loguearse", err)
	}

	return &user, nil
}

func (r *MainRepository) AuthUser(email string) (*models.User, error) {
	var user models.User

	err := r.DB.Preload("Role").
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "usuario no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener el usuario", err)
	}

	return &user, nil
}

func (r *MainRepository) AuthPointSale(userID uint, pointSaleID uint) (*models.PointSale, error) {
	var pointSale models.PointSale

	err := r.DB.Joins("JOIN user_point_sales ups ON ups.point_sale_id = point_sales.id").
		Where("ups.user_id = ?", userID).
		Find(&pointSale).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(403, "no existe relación del usuario con el punto de venta", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener la consulta", err)
	}

	return &pointSale, nil
}

func (r *MainRepository) LoginPointSale(userID uint, pointSaleID uint) (*models.PointSale, error) {
	var pointSale models.PointSale

	err := r.DB.Joins("JOIN user_point_sales ups ON ups.point_sale_id = point_sales.id").
		Where("ups.user_id = ?", userID).
		Find(&pointSale).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(403, "no existe relación del usuario con el punto de venta", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener la consulta", err)
	}

	return &pointSale, nil
}
