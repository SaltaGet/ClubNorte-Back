package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) IncomeGetByID(pointSaleID, id uint) (*models.Income, error) {
	var incomes *models.Income

	if err := r.DB.
		Preload("User").
		Preload("Items").
		Preload("Items.Product").
		First(&incomes, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "ingreso no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener los ingresos", err)
	}

	return incomes, nil
}

func (r *MainRepository) IncomeGetByDate(pointSaleID uint, fromDate, toDate time.Time, page, limit int) ([]*models.Income, int64, error) {
	var incomes []*models.Income

	offSet := (page - 1) * limit

	if err := r.DB.Preload("User").
		Where("created_at BETWEEN ? AND ?", fromDate, toDate).
		Where("point_sale_id = ?", pointSaleID).
		Order("created_at DESC").
		Offset(offSet).
		Limit(limit).
		Find(&incomes).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al obtener los ingresos", err)
	}

	var total int64
	if err := r.DB.Model(&models.Income{}).Where("created_at BETWEEN ? AND ?", fromDate, toDate).Count(&total).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al contar los ingresos", err)
	}

	return incomes, total, nil
}

// func (r *MainRepository) IncomeCreate(userID, pointSaleID uint, incomeCreate *schemas.IncomeCreate) (uint, error) {
// 	var incomeID uint
// 	err := r.DB.Transaction(func(tx *gorm.DB) error {
// 		var register models.Register
// 		if err := tx.
// 			Where("is_close = ? AND point_sale_id = ?", false, pointSaleID).
// 			Where("point_sale_id = ?", pointSaleID).
// 			Order("hour_open DESC").
// 			First(&register).Error; err != nil {
// 			if errors.Is(err, gorm.ErrRecordNotFound) {
// 				return schemas.ErrorResponse(400, "No se encontraron aperturas de caja, para registrar ingresos se debe de abrir la caja", err)
// 			}
// 			return schemas.ErrorResponse(500, "error al obtener la apertura de caja", err)
// 		}

// 		var incomeItems []*models.IncomeItem

// 		for _, item := range incomeCreate.Items {
// 			incomeItems = append(incomeItems, &models.IncomeItem{
// 				ProductID: item.ProductID,
// 				Quantity:  item.Quantity,
// 				Price:     item.Price,
// 				Subtotal:  item.Quantity * item.Price,
// 			})
// 		}

// 		total := 0.0
// 		for _, item := range incomeItems {
// 			total += item.Subtotal
// 		}

// 		var StockPointSale models.StockPointSale
// 		if err := tx.
// 			Where("point_sale_id = ?", pointSaleID).
// 			First(&StockPointSale).Error; err != nil {
// 			return schemas.ErrorResponse(500, "error al obtener el stock del punto de venta", err)
// 		}

// 		income := models.Income{
// 			PointSaleID:   pointSaleID,
// 			UserID:        userID,
// 			Total:         total,
// 			PaymentMethod: incomeCreate.PaymentMethod,
// 			Description:   incomeCreate.Description,
// 			RegisterID:    register.ID,
// 		}

// 		if err := tx.Create(&income).Error; err != nil {
// 			return schemas.ErrorResponse(500, "error al crear el ingreso", err)
// 		}

// 		incomeID = income.ID

// 		for i := range incomeItems {
// 			incomeItems[i].IncomeID = incomeID
// 		}

// 		if err := tx.Create(&incomeItems).Error; err != nil {
// 			return schemas.ErrorResponse(500, "error al crear los items del ingreso", err)
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		return 0, err
// 	}

// 	return incomeID, nil
// }

func (r *MainRepository) IncomeCreate(userID, pointSaleID uint, incomeCreate *schemas.IncomeCreate) (uint, error) {
	var incomeID uint
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		// 🔹 Buscar la caja abierta
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

		// 🔹 Armar items e ir validando stock
		var incomeItems []*models.IncomeItem
		total := 0.0

		for _, item := range incomeCreate.Items {
			var product models.Product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return schemas.ErrorResponse(400, fmt.Sprintf("El producto %d no existe", item.ProductID), err)
				}
				return schemas.ErrorResponse(500, "Error al obtener el producto", err)
			}
			// Buscar stock del producto en el punto de venta
			var stock models.StockPointSale
			if err := tx.
				Where("point_sale_id = ? AND product_id = ?", pointSaleID, item.ProductID).
				First(&stock).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return schemas.ErrorResponse(400, fmt.Sprintf("El producto %d no tiene stock en este punto de venta", item.ProductID), err)
				}
				return schemas.ErrorResponse(500, "Error al obtener stock", err)
			}

			// Validar stock suficiente
			if stock.Stock < float64(item.Quantity) {
				return schemas.ErrorResponse(
					400, 
					fmt.Sprintf("stock insuficiente para el producto %d (disponible: %.2f, requerido: %v)", item.ProductID, stock.Stock, item.Quantity), 
					fmt.Errorf("stock insuficiente para el producto %d (disponible: %.2f, requerido: %v)", item.ProductID, stock.Stock, item.Quantity),
				)
			}

			// Restar stock
			stock.Stock -= float64(item.Quantity)
			if err := tx.Save(&stock).Error; err != nil {
				return schemas.ErrorResponse(500, "Error al actualizar stock", err)
			}

			// Crear item en memoria
			incomeItems = append(incomeItems, &models.IncomeItem{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     product.Price,
				Subtotal:  item.Quantity * product.Price,
			})

			total += item.Quantity * product.Price
		}

		// 🔹 Crear ingreso
		income := models.Income{
			PointSaleID:   pointSaleID,
			UserID:        userID,
			Total:         total,
			PaymentMethod: incomeCreate.PaymentMethod,
			Description:   incomeCreate.Description,
			RegisterID:    register.ID,
		}

		if err := tx.Create(&income).Error; err != nil {
			return schemas.ErrorResponse(500, "Error al crear el ingreso", err)
		}
		incomeID = income.ID

		// 🔹 Asociar items
		for _, item := range incomeItems {
			item.IncomeID = incomeID
		}
		if err := tx.Create(&incomeItems).Error; err != nil {
			return schemas.ErrorResponse(500, "Error al crear items del ingreso", err)
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return incomeID, nil
}

// func (r *MainRepository) IncomeUpdate(userID, pointSaleID uint, incomeUpdate *schemas.IncomeUpdate) error {
// 	return r.DB.Transaction(func(tx *gorm.DB) error {
// 		// 1️⃣ Buscar el ingreso
// 		var income models.Income
// 		if err := tx.Preload("Items").First(&income, incomeUpdate.ID).Error; err != nil {
// 			if errors.Is(err, gorm.ErrRecordNotFound) {
// 				return schemas.ErrorResponse(404, "Ingreso no encontrado", err)
// 			}
// 			return schemas.ErrorResponse(500, "Error al obtener el ingreso", err)
// 		}

// 		// 2️⃣ Validar que pertenezca al usuario (opcional)
// 		// if income.UserID != userID {
// 		// 	return schemas.ErrorResponse(403, "No tiene permiso para actualizar este ingreso", nil)
// 		// }
// 		if income.PointSaleID != pointSaleID {
// 			return schemas.ErrorResponse(403, "El ingreso no pertenece al punto de venta", fmt.Errorf("el ingreso %d no pertenece al punto de venta %d", income.ID, pointSaleID))
// 		}

// 		// 3️⃣ Eliminar los items antiguos
// 		if len(income.Items) > 0 {
// 			if err := tx.Where("income_id = ?", income.ID).Delete(&models.IncomeItem{}).Error; err != nil {
// 				return schemas.ErrorResponse(500, "Error al eliminar los items antiguos", err)
// 			}
// 		}

// 		var newItems []*models.IncomeItem
// 		total := 0.0
// 		for _, item := range incomeUpdate.Items {
// 			subtotal := item.Quantity * item.Price
// 			newItems = append(newItems, &models.IncomeItem{
// 				IncomeID: income.ID,
// 				ProductID: item.ProductID,
// 				Quantity: item.Quantity,
// 				Price: item.Price,
// 				Subtotal: subtotal,
// 			})
// 			total += subtotal
// 		}
// 		if len(newItems) > 0 {
// 			if err := tx.Create(&newItems).Error; err != nil {
// 				return schemas.ErrorResponse(500, "Error al crear los nuevos items", err)
// 			}
// 		}

// 		// 5️⃣ Actualizar los campos del ingreso
// 		income.Total = total
// 		income.Description = incomeUpdate.Description
// 		income.PaymentMethod = incomeUpdate.PaymentMethod
// 		income.UserID = userID

// 		if err := tx.Save(&income).Error; err != nil {
// 			return schemas.ErrorResponse(500, "Error al actualizar el ingreso", err)
// 		}

// 		return nil
// 	})
// }
func (r *MainRepository) IncomeUpdate(userID, pointSaleID uint, incomeUpdate *schemas.IncomeUpdate) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// 1️⃣ Buscar el ingreso existente
		var income models.Income
		if err := tx.Preload("Items").First(&income, incomeUpdate.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return schemas.ErrorResponse(404, "Ingreso no encontrado", err)
			}
			return schemas.ErrorResponse(500, "Error al obtener el ingreso", err)
		}

		// 2️⃣ Validar que pertenezca al punto de venta
		if income.PointSaleID != pointSaleID {
			return schemas.ErrorResponse(403, "El ingreso no pertenece al punto de venta", 
				fmt.Errorf("el ingreso %d no pertenece al punto de venta %d", income.ID, pointSaleID))
		}

		// 3️⃣ Revertir stock de los ítems antiguos
		for _, oldItem := range income.Items {
			if err := tx.Model(&models.Product{}).
				Where("id = ?", oldItem.ProductID).
				Update("stock", gorm.Expr("stock + ?", oldItem.Quantity)).Error; err != nil {
				return schemas.ErrorResponse(500, "Error al revertir el stock de los productos", err)
			}
		}

		// 4️⃣ Eliminar los ítems antiguos
		if len(income.Items) > 0 {
			if err := tx.Where("income_id = ?", income.ID).Delete(&models.IncomeItem{}).Error; err != nil {
				return schemas.ErrorResponse(500, "Error al eliminar los items antiguos", err)
			}
		}

		// 5️⃣ Crear los nuevos ítems y ajustar stock
		var newItems []*models.IncomeItem
		total := 0.0
		for _, item := range incomeUpdate.Items {
			var product models.Product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return schemas.ErrorResponse(400, fmt.Sprintf("El producto %d no existe", item.ProductID), err)
				}
				return schemas.ErrorResponse(500, "Error al obtener el producto", err)
			}

			subtotal := item.Quantity * product.Price

			newItems = append(newItems, &models.IncomeItem{
				IncomeID:  income.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     product.Price,
				Subtotal:  subtotal,
			})
			total += subtotal

			// Descontar stock según la nueva cantidad
			if err := tx.Model(&models.Product{}).
				Where("id = ?", item.ProductID).
				Update("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
				return schemas.ErrorResponse(500, "Error al actualizar el stock de los productos", err)
			}
		}

		// 6️⃣ Guardar los nuevos ítems
		if len(newItems) > 0 {
			if err := tx.Create(&newItems).Error; err != nil {
				return schemas.ErrorResponse(500, "Error al crear los nuevos items", err)
			}
		}

		// 7️⃣ Actualizar el ingreso
		income.Total = total
		income.Description = incomeUpdate.Description
		income.PaymentMethod = incomeUpdate.PaymentMethod
		income.UserID = userID

		if err := tx.Save(&income).Error; err != nil {
			return schemas.ErrorResponse(500, "Error al actualizar el ingreso", err)
		}

		return nil
	})
}



func (r *MainRepository) IncomeDelete(pointSaleID, id uint) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// 1️⃣ Buscar el ingreso
		var income models.Income
		if err := tx.Preload("Items").
			Where("id = ? AND point_sale_id = ?", id, pointSaleID).
			First(&income).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return schemas.ErrorResponse(404, "Ingreso no encontrado", err)
			}
			return schemas.ErrorResponse(500, "Error al obtener el ingreso", err)
		}

		// 2️⃣ Eliminar items asociados
		if len(income.Items) > 0 {
			if err := tx.Where("income_id = ?", income.ID).Delete(&models.IncomeItem{}).Error; err != nil {
				return schemas.ErrorResponse(500, "Error al eliminar los items asociados", err)
			}
		}

		// 3️⃣ Eliminar el ingreso
		if err := tx.Delete(&income).Error; err != nil {
			return schemas.ErrorResponse(500, "Error al eliminar el ingreso", err)
		}

		return nil
	})
}

