package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) IncomeSportCourtGetByID(pointSaleID, id uint) (*models.IncomeSportsCourts, error) {
	var income models.IncomeSportsCourts
	if err := r.DB.Preload("SportsCourt").Preload("User").Where("id = ? AND point_sale_id = ?", id, pointSaleID).First(&income).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "ingreso no encontrado", err)
		}
		return nil, err
	}

	return &income, nil
}

func (r *MainRepository) IncomeSportCourtGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*models.IncomeSportsCourts, int64, error) {
	var incomes []*models.IncomeSportsCourts

	offSet := (page - 1) * limit

	if err := r.DB.Preload("SportsCourt").
		Where("created_at BETWEEN ? AND ?", fromDate, toDate).
		Where("point_sale_id = ?", pointSaleID).
		Order("created_at DESC").
		Offset(offSet).
		Limit(limit).
		Find(&incomes).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al obtener los ingresos", err)
	}

	var total int64
	if err := r.DB.Model(&models.IncomeSportsCourts{}).
		Where("created_at BETWEEN ? AND ?", fromDate, toDate).
		Where("point_sale_id = ?", pointSaleID).
		Count(&total).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al contar los ingresos", err)
	}

	return incomes, total, nil
}

func (r *MainRepository) IncomeSportCourtCreate(userID, pointSaleID uint, incomeCreate *schemas.IncomeSportsCourtsCreate) (uint, error) {
	// 🔹 Buscar la caja abierta
	var register models.Register
	if err := r.DB.
		Where("is_close = ? AND point_sale_id = ?", false, pointSaleID).
		Order("hour_open DESC").
		First(&register).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, schemas.ErrorResponse(400, "No hay caja abierta para este punto de venta", err)
		}
		return 0, schemas.ErrorResponse(500, "Error al obtener la apertura de caja", err)
	}

	loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
	now := time.Now().In(loc)

	income := models.IncomeSportsCourts{
		SportsCourtID:        incomeCreate.SportsCourtID,
		Shift:                incomeCreate.Shift,
		DatePlay:             incomeCreate.DatePlay,
		UserID:               userID,
		PartialPay:           incomeCreate.PartialPay,
		PartialPaymentMethod: incomeCreate.PartialPaymentMethod,
		PartialRegisterID:    pointSaleID,
		Total:                incomeCreate.Total,
		PointSaleID:          pointSaleID,
		DatePartialPay:       now,
	}

	if incomeCreate.Total == incomeCreate.PartialPay {
		restPay := 0.0
		income.RestPay = &restPay
		income.RestPaymentMethod = &incomeCreate.PartialPaymentMethod
		income.RestRegisterID = &pointSaleID
	}

	if err := r.DB.Create(&income).Error; err != nil {
		return 0, schemas.ErrorResponse(500, "Error al crear el ingreso", err)
	}

	return income.ID, nil
}

// func (r *MainRepository) IncomeSportCourtUpdate(userID, pointSaleID uint, incomeUpdate *schemas.IncomeSportsCourtsUpdate) error {
// 	return r.DB.Transaction(func(tx *gorm.DB) error {
// 		updateIncome := map[string]any{
// 			"sport_court_id": incomeUpdate.SportsCourtID,

// 			"shift":                  incomeUpdate.Shift,
// 			"date_play":              incomeUpdate.DatePlay,
// 			"partial_pay":            incomeUpdate.PartialPay,
// 			"partial_payment_method": incomeUpdate.PartialPaymentMethod,
// 			"date_partial_pay":       incomeUpdate.DatePartialPay,

// 			"rest_pay":            incomeUpdate.RestPay,
// 			"rest_payment_method": incomeUpdate.RestPaymentMethod,
// 			"date_rest_pay":       incomeUpdate.DateRestPay,

// 			"price": incomeUpdate.Price,
// 		}

// 		if err := tx.Model(&models.IncomeSportsCourts{}).Where("id = ? AND point_sale_id = ?", incomeUpdate.ID, pointSaleID).Updates(updateIncome).Error; err != nil {
// 			if errors.Is(err, gorm.ErrRecordNotFound) {
// 				return schemas.ErrorResponse(404, "Ingreso no encontrado", err)
// 			}
// 			return schemas.ErrorResponse(500, "Error al obtener el ingreso", err)
// 		}

// 		if tx.RowsAffected == 0 {
// 			return schemas.ErrorResponse(404, "Ingreso no encontrado", nil)
// 		}

// 		return nil
// 	})
// }

func (r *MainRepository) IncomeSportCourtUpdatePay(userID, pointSaleID uint, incomeUpdate *schemas.IncomeSportsCourtsRestPay) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var income models.IncomeSportsCourts
		if err := tx.Where("id = ? AND point_sale_id = ?", incomeUpdate.ID, pointSaleID).First(&income).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return schemas.ErrorResponse(404, "Ingreso no encontrado", err)
			}
			return schemas.ErrorResponse(500, "Error al obtener el ingreso", err)
		}


		now := time.Now()
		restPay := income.Total - income.PartialPay

		if restPay != incomeUpdate.RestPay {
			return schemas.ErrorResponse(400, "Error al actualizar el ingreso", fmt.Errorf("el pago restante debe ser igual al precio menos el pago parcial"))
		}

		income.RestPay = &incomeUpdate.RestPay
		income.RestPaymentMethod = &incomeUpdate.RestPaymentMethod
		income.DateRestPay = &now

		if tx.Save(&income).Error != nil {
			return schemas.ErrorResponse(500, "error al actualizar ingreso", fmt.Errorf("error al actualizar el ingreso"))
		}

		return nil
	})
}

func (r *MainRepository) IncomeSportCourtDelete(pointSaleID, id uint) error {
	if err := r.DB.Where("id = ? AND point_sale_id = ?", id, pointSaleID).Delete(&models.IncomeSportsCourts{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schemas.ErrorResponse(404, "Ingreso no encontrado", err)
		}
		return schemas.ErrorResponse(500, "Error al eliminar el ingreso", err)
	}

	return nil
}
