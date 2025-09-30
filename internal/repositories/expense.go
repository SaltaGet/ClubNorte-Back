package repositories

import (
	"errors"
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) ExpenseGetByID(pointSaleID, id uint) (*models.Expense, error) {
	var expense *models.Expense

	if err := r.DB.
		Preload("User").
		First(&expense, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "egreso no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener egreso", err)
	}

	return expense, nil
}

func (r *MainRepository) ExpenseGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*models.Expense, int64, error) {
	var expenses []*models.Expense

	offSet := (page - 1) * limit

	if err := r.DB.Preload("User").
		Where("created_at BETWEEN ? AND ?", fromDate, toDate).
		Where("point_sale_id = ?", pointSaleID).
		Order("created_at DESC").
		Offset(offSet).
		Limit(limit).
		Find(&expenses).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al obtener los ingresos", err)
	}

	var total int64
	if err := r.DB.Model(&models.Expense{}).
		Where("created_at BETWEEN ? AND ?", fromDate, toDate).
		Where("point_sale_id = ?", pointSaleID).
		Count(&total).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al contar los ingresos", err)
	}

	return expenses, total, nil
}

func (r *MainRepository) ExpenseCreate(userID, pointSaleID uint, expenseCreate *schemas.ExpenseCreate) (uint, error) {
	var expenseID uint
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		var register models.Register
		if err := tx.
			Where("is_close = ? AND point_sale_id = ?", false, pointSaleID).
			Order("hour_open DESC").
			First(&register).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return schemas.ErrorResponse(400, "No hay caja abierta para este punto de venta", err)
			}
			return schemas.ErrorResponse(500, "Error al obtener la apertura de caja", err)
		}

		expense := models.Expense{
			PointSaleID:   pointSaleID,
			UserID:        userID,
			Total:         expenseCreate.Total,
			PaymentMethod: expenseCreate.PaymentMethod,
			Description:   expenseCreate.Description,
			RegisterID:    register.ID,
		}

		if err := tx.Create(&expense).Error; err != nil {
			return schemas.ErrorResponse(500, "Error al crear el egreso", err)
		}
		expenseID = expense.ID

		return nil
	})

	if err != nil {
		return 0, err
	}

	return expenseID, nil
}

func (r *MainRepository) ExpenseDelete(pointSaleID, id uint) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var expense models.Expense
		if err := tx.Where("id = ? AND point_sale_id = ?", id, pointSaleID).
			First(&expense).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return schemas.ErrorResponse(404, "Ingreso no encontrado", err)
			}
			return schemas.ErrorResponse(500, "Error al obtener el ingreso", err)
		}

		if err := tx.Delete(&expense).Error; err != nil {
			return schemas.ErrorResponse(500, "Error al eliminar el ingreso", err)
		}

		return nil
	})
}


