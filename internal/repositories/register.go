package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) RegisterExistOpen(pointSaleID uint) (bool, error) {
	var existRegisterOpen float64
	if err := r.DB.
		Model(&models.Register{}).
		Select("count(*)").
		Where("is_close = ? AND point_sale_id = ?", false, pointSaleID).
		Scan(&existRegisterOpen).Error; err != nil {
		return false, schemas.ErrorResponse(500, "error al contar aperturas de caja", err)
	}

	if existRegisterOpen > 0 {
		return true, nil
	}

	return false, nil
}

func (r *MainRepository) RegisterOpen(pointSaleID uint, userID uint, amountOpen schemas.RegisterOpen) error {
	var existRegisterOpen float64

	if err := r.DB.
		Model(&models.Register{}).
		Select("count(*)").
		Where("is_close = ? AND point_sale_id = ?", false, pointSaleID).
		Scan(&existRegisterOpen).Error; err != nil {
		return schemas.ErrorResponse(500, "error al contar aperturas de caja", err)
	}

	if existRegisterOpen > 0 {
		return schemas.ErrorResponse(400, "ya existe una apertura de caja, antes de continuar cierre la caja", fmt.Errorf("ya existe una apertura de caja, antes de continuar cerrar"))
	}

	registerOpen := models.Register{
		PointSaleID: pointSaleID,
		UserOpenID:  userID,
		OpenAmount:  amountOpen.OpenAmount,
		HourOpen:    time.Now().UTC(),
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

	var user models.User
	if err := r.DB.Preload("Role").First(&user, userID).Error; err != nil {
		return schemas.ErrorResponse(404, "usuario no encontrado", err)
	}

	if user.Role.Name != "admin" || user.ID != register.UserOpenID {
		return schemas.ErrorResponse(403, "no tienes permiso para cerrar la caja, solo el creador o el admin", fmt.Errorf("no tienes permiso para cerrar la caja, solo el creador o el admin"))
	}

	// var totalsIncome schemas.Totals
	// if err := r.DB.
	// 	Model(&models.Income{}).
	// 	Select(`
	// 	COALESCE(SUM(CASE WHEN payment_method = 'efectivo' THEN COALESCE(total,0) ELSE 0	END),0) AS cash,
	// 	COALESCE(SUM(CASE WHEN payment_method IN ('tarjeta','transferencia') THEN COALESCE(total,0) ELSE 0 END),0) AS others
	// `).
	// 	Where("register_id = ?", register.ID).
	// 	Scan(&totalsIncome).Error; err != nil {
	// 	return schemas.ErrorResponse(500, "error al obtener ingresos por métodos", err)
	// }

	// var totalsIncomeCourts schemas.Totals
	// if err := r.DB.
	// 	Model(&models.IncomeSportsCourts{}).
	// 	Select(`
	// 	COALESCE(
	// 		SUM(
	// 			CASE WHEN partial_payment_method = 'efectivo' THEN COALESCE(partial_pay,0) ELSE 0	END + 
	// 			CASE WHEN rest_payment_method = 'efectivo' THEN COALESCE(rest_pay,0) ELSE 0 END
	// 		),0
	// 	) AS cash,
	// 	COALESCE(
  //       SUM(
  //           CASE WHEN partial_payment_method IN ('tarjeta','transferencia') THEN COALESCE(partial_pay,0) ELSE 0 END +
  //           CASE WHEN rest_payment_method IN ('tarjeta','transferencia') THEN COALESCE(rest_pay,0) ELSE 0 END
  //       ), 0
  //   ) AS others
	// `).
	// 	Where("partial_register_id = ? OR rest_register_id = ?", register.ID, register.ID).
	// 	Scan(&totalsIncomeCourts).Error; err != nil {
	// 	return schemas.ErrorResponse(500, "error al obtener ingresos de canchas por métodos", err)
	// }

	// var totalsExpense schemas.Totals
	// if err := r.DB.
	// 	Model(&models.Expense{}).
	// 	Select(`
	// 	COALESCE(SUM(CASE WHEN payment_method = 'efectivo' THEN COALESCE(total,0) ELSE 0	END),0) AS cash,
	// 	COALESCE(SUM(CASE WHEN payment_method IN ('tarjeta','transferencia') THEN COALESCE(total,0) ELSE 0 END),0) AS others
	// `).
	// 	Where("register_id = ?", register.ID).
	// 	Scan(&totalsExpense).Error; err != nil {
	// 	return schemas.ErrorResponse(500, "error al obtener ingresos por métodos", err)
	// }

	now := time.Now().UTC()
	// totalIncomesCash := totalsIncome.Cash + totalsIncomeCourts.Cash
	// totalIncomesOthers := totalsIncome.Others + totalsIncomeCourts.Others
	// totalExpenseCash := totalsExpense.Cash + totalsExpense.Cash
	// totalExpenseOthers := totalsExpense.Others + totalsExpense.Others

	register.CloseAmount = &amountOpen.CloseAmount
	register.IsClose = true
	register.HourClose = &now
	register.UserCloseID = &userID
	// register.TotalIncomeCash = &totalIncomesCash
	// register.TotalIncomeOthers = &totalIncomesOthers
	// register.TotalExpenseCash = &totalExpenseCash
	// register.TotalExpenseOthers = &totalExpenseOthers

	if err := r.DB.Save(&register).Error; err != nil {
		return schemas.ErrorResponse(500, "error al cerrar la caja", err)
	}

	return nil
}

func (r *MainRepository) RegisterInform(pointSaleID uint, userID uint, fromDate, toDate time.Time) ([]*models.Register, error) {
	var registers []*models.Register
	if err := r.DB.
		Preload("UserOpen").
		Preload("UserClose").
		// Preload("PointSale").
		Where("point_sale_id = ? AND created_at >= ? AND created_at <= ?", pointSaleID, fromDate, toDate).
		Order("created_at DESC").
		Find(&registers).Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(registers); i++ {
		var totalsIncome schemas.Totals
		if err := r.DB.
			Model(&models.Income{}).
			Select(`
				COALESCE(SUM(CASE WHEN payment_method = 'efectivo' THEN COALESCE(total,0) ELSE 0	END),0) AS cash,
				COALESCE(SUM(CASE WHEN payment_method IN ('tarjeta','transferencia') THEN COALESCE(total,0) ELSE 0 END),0) AS others
			`).
			Where("register_id = ?", registers[i].ID).
			Scan(&totalsIncome).Error; err != nil {
			return nil, schemas.ErrorResponse(500, "error al obtener ingresos por métodos", err)
		}

		var totalsExpense schemas.Totals
		if err := r.DB.
			Model(&models.Expense{}).
			Select(`
				COALESCE(SUM(CASE WHEN payment_method = 'efectivo' THEN COALESCE(total,0) ELSE 0	END),0) AS cash,
				COALESCE(SUM(CASE WHEN payment_method IN ('tarjeta','transferencia') THEN COALESCE(total,0) ELSE 0 END),0) AS others
			`).
			Where("register_id = ?", registers[i].ID).
			Scan(&totalsExpense).Error; err != nil {
			return nil, schemas.ErrorResponse(500, "error al obtener ingresos por métodos", err)
		}

		var totalsIncomeCourts schemas.Totals
		if err := r.DB.
			Model(&models.IncomeSportsCourts{}).
			Select(`
				COALESCE(
					SUM(
						CASE WHEN partial_payment_method = 'efectivo' THEN COALESCE(partial_pay,0) ELSE 0	END + 
						CASE WHEN rest_payment_method = 'efectivo' THEN COALESCE(rest_pay,0) ELSE 0 END
					),0
				) AS cash,
				COALESCE(
						SUM(
								CASE WHEN partial_payment_method IN ('tarjeta','transferencia') THEN COALESCE(partial_pay,0) ELSE 0 END +
								CASE WHEN rest_payment_method IN ('tarjeta','transferencia') THEN COALESCE(rest_pay,0) ELSE 0 END
						), 0
				) AS others
			`).
			Where("partial_register_id = ? OR rest_register_id = ?", registers[i].ID, registers[i].ID).
			Scan(&totalsIncomeCourts).Error; err != nil {
			return nil, schemas.ErrorResponse(500, "error al obtener ingresos de canchas por métodos", err)
		}

		totalIncomesCash := totalsIncome.Cash + totalsIncomeCourts.Cash
		totalIncomesOthers := totalsIncome.Others + totalsIncomeCourts.Others
		totalExpenseCash := totalsExpense.Cash + totalsExpense.Cash
		totalExpenseOthers := totalsExpense.Others + totalsExpense.Others

		registers[i].TotalIncomeCash = &totalIncomesCash
		registers[i].TotalIncomeOthers = &totalIncomesOthers
		registers[i].TotalExpenseCash = &totalExpenseCash
		registers[i].TotalExpenseOthers = &totalExpenseOthers
	}

	return registers, nil
}

