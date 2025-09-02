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
	if err := r.DB.Preload("User").Preload("Product").First(&movement, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("movimiento no encontrado")
		}
		return nil, err
	}
	return movement, nil
}

func (r *MainRepository) MovementStockGetAll(page, limit int) ([]*models.MovementStock, int64, error) {
	offset := (page - 1) * limit

	var movements []*models.MovementStock
	var total int64
	if err := r.DB.Preload("User").Preload("Product").Offset(offset).Limit(limit).Find(&movements).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Model(&models.MovementStock{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return movements, total, nil
}

func (r *MainRepository) MoveStock(userID uint, input *schemas.MovementStock) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var fromID, toID uint

		switch input.FromType {
		case "deposit":
			var deposit models.StockDeposit
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("product_id = ?", input.ProductID).
				FirstOrCreate(&deposit, &models.StockDeposit{ProductID: input.ProductID}).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf("producto no encontrado en depósito origen")
				}
				return err
			}
			fromID = deposit.ID

			if !input.IgnoreStock && deposit.Stock < input.Amount {
				return fmt.Errorf("no hay suficiente stock en depósito para transferir %.2f unidades", input.Amount)
			}

			deposit.Stock -= input.Amount

			defer func() {
				_ = tx.Save(&deposit).Error
			}()

		case "point_sale":
			var ps models.StockPointSale
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("product_id = ? AND point_sale_id = ?", input.ProductID, input.FromID).
				FirstOrCreate(&ps, models.StockPointSale{
					ProductID:   input.ProductID,
					PointSaleID: input.FromID,
				}).Error; err != nil {
				return err
			}
			fromID = ps.ID

			if !input.IgnoreStock && ps.Stock < input.Amount {
				return fmt.Errorf("no hay suficiente stock en punto de venta para transferir %.2f unidades", input.Amount)
			}

			ps.Stock -= input.Amount

			defer func() {
				_ = tx.Save(&ps).Error
			}()
		default:
			return fmt.Errorf("tipo de origen inválido: %s", input.FromType)
		}

		switch input.ToType {
		case "deposit":
			var deposit models.StockDeposit
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("product_id = ?", input.ProductID).
				First(&deposit).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf("producto no encontrado en depósito destino")
				}
				return err
			}
			deposit.Stock += input.Amount
			toID = deposit.ID

			defer func() {
				_ = tx.Save(&deposit).Error
			}()

		case "point_sale":
			var ps models.StockPointSale
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("product_id = ? AND point_sale_id = ?", input.ProductID, input.ToID).
				FirstOrCreate(&ps, models.StockPointSale{
					ProductID:   input.ProductID,
					PointSaleID: input.ToID,
				}).Error; err != nil {
				return err
			}
			ps.Stock += input.Amount
			toID = ps.ID

			defer func() {
				_ = tx.Save(&ps).Error
			}()
		default:
			return fmt.Errorf("tipo de destino inválido: %s", input.ToType)
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
			return err
		}

		return nil
	})
}
