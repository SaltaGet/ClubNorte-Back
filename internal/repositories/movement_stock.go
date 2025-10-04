package repositories

import (
	"errors"
	"fmt"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *MainRepository) MovementStockGetByID(id uint) (*models.MovementStock, error) {
	var movement *models.MovementStock
	if err := r.DB.Preload("User").Preload("Product").Preload("Product.Category").First(&movement, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "movimiento no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener el movimiento", err)
	}
	return movement, nil
}

func (r *MainRepository) MovementStockGetAll(page, limit int) ([]*models.MovementStock, int64, error) {
	offset := (page - 1) * limit

	var movements []*models.MovementStock
	var total int64
	if err := r.DB.
		Preload("User").
		Preload("Product").
		Offset(offset).
		Limit(limit).
		Order("created_at desc").
		Find(&movements).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al obtener movimientos", err)
	}

	if err := r.DB.Model(&models.MovementStock{}).Count(&total).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al contar movimientos", err)
	}

	return movements, total, nil
}

func (r *MainRepository) MoveStock(userID uint, input *schemas.MovementStock) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var fromID, toID uint

		switch input.FromType {
		case "deposit":
			var product models.Product
			if err := tx.First(&product, input.ProductID).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return schemas.ErrorResponse(404, "producto no encontrado", err)
				}
				return schemas.ErrorResponse(500, "error al obtener el producto", err)
			}

			if product.Price <= 0.0 {
				return schemas.ErrorResponse(400, "No se puede editar un producto sin precio", fmt.Errorf("eNo se puede editar un producto sin precio"))
			}

			var deposit models.StockDeposit
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("product_id = ?", input.ProductID).
				FirstOrCreate(&deposit, &models.StockDeposit{ProductID: input.ProductID}).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return schemas.ErrorResponse(404, "deposito no encontrado", err)
				}
				return schemas.ErrorResponse(500, "error al obtener el depósito", err)
			}
			fromID = 100

			if !input.IgnoreStock && deposit.Stock < input.Amount {
				return schemas.ErrorResponse(400, "no hay suficiente stock en depósito para transferir", fmt.Errorf("no hay suficiente stock en depósito para transferir: %.2f", input.Amount))
			}

			deposit.Stock -= input.Amount

			defer func() {
				_ = tx.Save(&deposit).Error
			}()

		case "point_sale":
			var pointSaleExist bool
			if err := tx.Model(&models.PointSale{}).
				Select("count(*) > 0").
				Where("id = ?", input.FromID).
				Find(&pointSaleExist).Error; err != nil {
				return schemas.ErrorResponse(500, "error al obtener el punto de venta", err)
			}
			
			if !pointSaleExist {
				return schemas.ErrorResponse(404, "punto de venta no encontrado", fmt.Errorf("punto de venta no encontrado"))
			}

			var ps models.StockPointSale
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("product_id = ? AND point_sale_id = ?", input.ProductID, input.FromID).
				FirstOrCreate(&ps, models.StockPointSale{
					ProductID:   input.ProductID,
					PointSaleID: input.FromID,
				}).Error; err != nil {
				return schemas.ErrorResponse(500, "error al obtener producto del punto de venta", err)
			}
			fromID = input.FromID

			if !input.IgnoreStock && ps.Stock < input.Amount {
				return schemas.ErrorResponse(400, "no hay suficiente stock en punto de venta para transferir", fmt.Errorf("no hay suficiente stock en punto de venta para transferir %.2f unidades", input.Amount))
			}

			ps.Stock -= input.Amount

			defer func() {
				_ = tx.Save(&ps).Error
			}()
		default:
			return schemas.ErrorResponse(400, "tipo de origen inválido", fmt.Errorf("tipo de origen inválido: %s", input.FromType))
		}

		switch input.ToType {
		case "deposit":
			var deposit models.StockDeposit
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("product_id = ?", input.ProductID).
				First(&deposit).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return schemas.ErrorResponse(404, "producto no encontrado en depósito destino", err)
				}
				return schemas.ErrorResponse(500, "error al obtener el depósito de destino", err)
			}
			deposit.Stock += input.Amount
			toID = 100

			defer func() {
				_ = tx.Save(&deposit).Error
			}()

		case "point_sale":
			var pointSaleExist bool
			if err := tx.Model(&models.PointSale{}).
				Select("count(*) > 0").
				Where("id = ?", input.ToID).
				Find(&pointSaleExist).Error; err != nil {
				return schemas.ErrorResponse(500, "error al obtener el punto de venta", err)
			}
			
			if !pointSaleExist {
				return schemas.ErrorResponse(404, "punto de venta no encontrado", fmt.Errorf("punto de venta no encontrado"))
			}

			var ps models.StockPointSale
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("product_id = ? AND point_sale_id = ?", input.ProductID, input.ToID).
				FirstOrCreate(&ps, models.StockPointSale{
					ProductID:   input.ProductID,
					PointSaleID: input.ToID,
				}).Error; err != nil {
				return schemas.ErrorResponse(500, "error al obtener el punto de venta de destino", err)
			}
			ps.Stock += input.Amount
			toID = input.ToID

			defer func() {
				_ = tx.Save(&ps).Error
			}()
		default:
			return schemas.ErrorResponse(400, "tipo de destino inválido", fmt.Errorf("tipo de destino inválido: %s", input.ToType))
		}

		movementStock := models.MovementStock{
			UserID:      userID,
			ProductID:   input.ProductID,
			Amount:      input.Amount,
			FromID:      fromID,
			FromType:    input.FromType,
			ToID:        toID,
			ToType:      input.ToType,
			IgnoreStock: input.IgnoreStock,
		}

		if err := tx.Create(&movementStock).Error; err != nil {
			return schemas.ErrorResponse(500, "error al realizar el movimiento", err)
		}

		return nil
	})
}
