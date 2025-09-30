package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) ExpenseBuyGetByID(id uint) (*models.ExpenseBuy, error) {
	var expenseBuy *models.ExpenseBuy

	if err := r.DB.
		Preload("User").
		Preload("ItemExpenseBuys").
		Preload("ItemExpenseBuys.Product").
		First(&expenseBuy, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "egreso de compras no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener los egreso de compras", err)
	}

	// var total float64
	// for _, item := range expenseBuy.ItemExpenseBuys {
	// 	total += item.Subtotal
	// }
	// expenseBuy.Total = total

	return expenseBuy, nil
}

func (r *MainRepository) ExpenseBuyGetByDate(fromDate, toDate time.Time, page, limit int) ([]*models.ExpenseBuy, int64, error) {
	var expensesBuy []*models.ExpenseBuy

	offSet := (page - 1) * limit

	if err := r.DB.Preload("User").
		Where("created_at BETWEEN ? AND ?", fromDate, toDate).
		Order("created_at DESC").
		Offset(offSet).
		Limit(limit).
		Find(&expensesBuy).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al obtener los egresos de compras", err)
	}

	var total int64
	if err := r.DB.Model(&models.ExpenseBuy{}).
		Where("created_at BETWEEN ? AND ?", fromDate, toDate).
		Count(&total).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al contar los egresos de compras", err)
	}

	// for _, expenseBuy := range expensesBuy {
	// 	var total float64
	// 	for _, item := range expenseBuy.ItemExpenseBuys {
	// 		total += item.Subtotal
	// 	}
	// 	expenseBuy.Total = total
	// }

	return expensesBuy, total, nil
}

func (r *MainRepository) ExpenseBuyCreate(userID uint, expenseBuyCreate *schemas.ExpenseBuyCreate) (uint, error) {
	var expenseBuyID uint
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		// 🔹 Buscar la caja abierta
		// var register models.Register
		// if err := tx.
		// 	Where("is_close = ? AND point_sale_id = ?", false, pointSaleID).
		// 	Order("hour_open DESC").
		// 	First(&register).Error; err != nil {
		// 	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 		return schemas.ErrorResponse(400, "No hay caja abierta para este punto de venta", err)
		// 	}
		// 	return schemas.ErrorResponse(500, "Error al obtener la apertura de caja", err)
		// }

		// 🔹 Armar items e ir validando stock
		var expenseItems []*models.ItemExpenseBuy
		total := 0.0

		for _, item := range expenseBuyCreate.ItemExpenseBuys {
			if item.Quantity <= 0 {
				return schemas.ErrorResponse(400, fmt.Sprintf("La cantidad para el producto %d no es válida", item.ProductID), fmt.Errorf("la cantidad para el producto %d no es válida", item.ProductID))
			}

			var product models.Product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return schemas.ErrorResponse(400, fmt.Sprintf("El producto %d no existe", item.ProductID), err)
				}
				return schemas.ErrorResponse(500, "Error al obtener el producto", err)
			}
			// Buscar stock del producto en el punto de venta
			var stock models.StockDeposit
			if err := tx.
				Where("product_id = ?", item.ProductID).
				FirstOrCreate(&stock, models.StockDeposit{ProductID: item.ProductID, Stock: 0}).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return schemas.ErrorResponse(400, fmt.Sprintf("El producto %d no tiene stock en este punto de venta", item.ProductID), err)
				}
				return schemas.ErrorResponse(500, "Error al obtener stock", err)
			}

			if err := tx.Model(&stock).
				Update("stock", gorm.Expr("stock + ?", item.Quantity)).Error; err != nil {
				return schemas.ErrorResponse(500, "Error al actualizar stock", err)
			}
			// stock.Stock += item.Quantity
			// if err := tx.Save(&stock).Error; err != nil {
			// 	return schemas.ErrorResponse(500, "Error al actualizar stock", err)
			// }

			subTotal := item.Quantity * item.Price

			expenseItems = append(expenseItems, &models.ItemExpenseBuy{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Price,
				Subtotal:  subTotal,
			})

			total += subTotal
		}

		// 🔹 Crear ingreso
		expenseBuy := models.ExpenseBuy{
			UserID:        userID,
			PaymentMethod: expenseBuyCreate.PaymentMethod,
			Description:   expenseBuyCreate.Description,
			Total:         total,
		}

		if err := tx.Create(&expenseBuy).Error; err != nil {
			return schemas.ErrorResponse(500, "Error al crear el ingreso", err)
		}
		expenseBuyID = expenseBuy.ID

		// 🔹 Asociar items
		for _, item := range expenseItems {
			item.ExpenseBuyID = expenseBuyID
		}
		if err := tx.Create(&expenseItems).Error; err != nil {
			return schemas.ErrorResponse(500, "Error al crear items del ingreso", err)
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return expenseBuyID, nil
}

func (r *MainRepository) ExpenseBuyDelete(id uint) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// 1️⃣ Buscar la compra con sus items
		var expenseBuy models.ExpenseBuy
		if err := tx.Preload("ItemExpenseBuys").
			Where("id = ?", id).
			First(&expenseBuy).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return schemas.ErrorResponse(404, "Compra no encontrada", err)
			}
			return schemas.ErrorResponse(500, "Error al obtener la compra", err)
		}

		// 2️⃣ Restar stock de cada item
		for _, item := range expenseBuy.ItemExpenseBuys {
			var stock models.StockDeposit
			if err := tx.
				Where("product_id = ?", item.ProductID).
				First(&stock).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return schemas.ErrorResponse(400, fmt.Sprintf("No existe stock para el producto %d en este punto de venta", item.ProductID), err)
				}
				return schemas.ErrorResponse(500, "Error al obtener stock del producto", err)
			}

			if stock.Stock < item.Quantity {
				return schemas.ErrorResponse(400, fmt.Sprintf("Stock insuficiente para revertir el producto %d", item.ProductID), nil)
			}

			if err := tx.Model(&stock).
				Update("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
				return schemas.ErrorResponse(500, "Error al actualizar stock del producto", err)
			}
		}

		if len(expenseBuy.ItemExpenseBuys) > 0 {
			if err := tx.Where("expense_buy_id = ?", expenseBuy.ID).
				Delete(&models.ItemExpenseBuy{}).Error; err != nil {
				return schemas.ErrorResponse(500, "Error al eliminar los items asociados", err)
			}
		}

		if err := tx.Delete(&expenseBuy).Error; err != nil {
			return schemas.ErrorResponse(500, "Error al eliminar la compra", err)
		}

		return nil
	})
}
