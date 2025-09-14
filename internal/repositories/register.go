package repositories

import (
	"errors"
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) RegisterOpen(pointSaleID uint, userID uint, amountOpen schemas.RegisterOpen) error {
	registerOpen := models.Register{
		PointSaleID: pointSaleID,
		UserID:      userID,
		OpenAmount:  amountOpen.OpenAmount,
	}

	if err := r.DB.Create(&registerOpen).Error; err != nil {
		return schemas.ErrorResponse(500, "error al registrar la apertura de caja", err)
	}

	return nil
}

func (r *MainRepository) RegisterClose(pointSaleID uint, userID uint, amountOpen schemas.RegisterClose) error {
	var register models.Register

	if err := r.DB.
		Where("is_close = ? AND point_sale_id = ?", false, pointSaleID).
		Order("hour_open DESC").
		First(&register).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schemas.ErrorResponse(404, "No se encontraron aperturas de caja", err)
		}
		return schemas.ErrorResponse(500, "error al obtener la apertura de caja", err)
	}

	var totalIncome float64
	if err := r.DB.
		Model(&models.Income{}).
		Select("SUM(amount)").
		Where("register_id = ?", register.ID).
		Scan(&totalIncome).Error; err != nil {
		return schemas.ErrorResponse(500, "error al obtener el total de ingresos", err)
	}

	now := time.Now().UTC()

	var totalIncomeCourts float64
	if err := r.DB.
		Model(&models.IncomeSportsCourts{}).
		Select("SUM(amount)").
		Where().
		Scan(&totalIncomeCourts).Error; err != nil {
		return schemas.ErrorResponse(500, "error al obtener el total de ingresos", err)
	}
Select("SUM(total_pay - COALESCE(partial_pay, 0))")


	register.CloseAmount = &amountOpen.CloseAmount
	register.IsClose = true
	register.HourClose = &now

	return nil
}

func (r *MainRepository) RegisterInform(pointSaleID uint, userID uint, dateInform time.Time) (*schemas.RegisterInform, error) {
	return nil, nil
}
