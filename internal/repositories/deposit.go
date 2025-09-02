package repositories

import (
	"errors"
	"fmt"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"gorm.io/gorm"
)

func (r *MainRepository) DepositGetByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.DB.Preload("Category").Preload("StockDeposit").Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *MainRepository) DepositGetByCode(code string) (*models.Product, error) {
	var product models.Product
	if err := r.DB.Preload("Category").Preload("StockDeposit").Where("code = ?", code).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *MainRepository) DepositGetByName(name string) ([]*models.Product, error) {
	var products []*models.Product
	if err := r.DB.Preload("Category").Preload("StockDeposit").Where("name LIKE", "%"+name+"%").First(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *MainRepository) DepositGetAll(page, limit int) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64
	if err := r.DB.Preload("Category").Preload("StockDeposit").Offset((page - 1) * limit).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}
	if err := r.DB.Model(&models.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func (r *MainRepository) DepositUpdateStock(productID uint, stock float64, method string) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var deposit models.StockDeposit

		if err := tx.Where("product_id = ?", productID).FirstOrCreate(&deposit, &models.StockDeposit{ProductID: productID}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("no existe un depósito para el producto %d", productID)
			}
			return err
		}

		switch method {
		case "add":
			deposit.Stock += stock
		case "subtract":
			if deposit.Stock < stock {
				return fmt.Errorf("no hay suficiente stock para restar %.2f unidades", stock)
			}
			deposit.Stock -= stock
		case "set":
			deposit.Stock = stock
		default:
			return fmt.Errorf("método inválido: %s", method)
		}

		if err := tx.Save(&deposit).Error; err != nil {
			return err
		}

		return nil
	})
}
