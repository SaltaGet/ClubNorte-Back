package repositories

import (
	"errors"
	"fmt"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) PointSaleGet(id uint) (*models.PointSale, error) {
	var pointSales models.PointSale

	if err := r.DB.Where("id = ?", id).First(&pointSales).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("punto de venta no encontrado")
		}
		return nil, err
	}

	return &pointSales, nil
}

func (r *MainRepository) PointSaleGetAll() ([]*models.PointSale, error) {
	var pointSales []*models.PointSale

	if err := r.DB.Find(&pointSales).Error; err != nil {
		return nil, err
	}

	return pointSales, nil
}

func (r *MainRepository) PointSaleCreate(pointSaleCreate *schemas.PointSaleCreate) (uint, error) {
	var pointSale models.PointSale

	pointSale.Name = pointSaleCreate.Name
	pointSale.Description = pointSaleCreate.Description

	if err := r.DB.Create(&pointSale).Error; err != nil {
		return 0, err
	}

	return pointSale.ID, nil
}

func (r *MainRepository) PointSaleUpdate(pointSaleUpdate *schemas.PointSaleUpdate) error {
	err := r.DB.Model(&models.PointSale{}).
		Where("id = ?", pointSaleUpdate.ID).
		Updates(models.PointSale{Name: pointSaleUpdate.Name, Description: pointSaleUpdate.Description}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("punto de venta no encontrado")
		}
		return err
	}

	return nil
}

func (r *MainRepository) PointSaleDelete(id uint) error {
	err := r.DB.Where("id = ?", id).Delete(&models.PointSale{}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("punto de venta no encontrado")
		}
		return err
	}

	return nil
}
