package repositories

import (
	"errors"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) PointSaleGet(id uint) (*models.PointSale, error) {
	var pointSales models.PointSale

	if err := r.DB.Where("id = ?", id).First(&pointSales).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "punto de venta no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener el punto de venta", err)
	}

	return &pointSales, nil
}

func (r *MainRepository) PointSaleGetAll() ([]*models.PointSale, error) {
	var pointSales []*models.PointSale

	if err := r.DB.Find(&pointSales).Error; err != nil {
		return nil, schemas.ErrorResponse(500, "error al obtener los puntos de venta", err)
	}

	return pointSales, nil
}

func (r *MainRepository) PointSaleCreate(pointSaleCreate *schemas.PointSaleCreate) (uint, error) {
	var pointSale models.PointSale

	pointSale.Name = pointSaleCreate.Name
	pointSale.Description = pointSaleCreate.Description

	if err := r.DB.Create(&pointSale).Error; err != nil {
		if IsDuplicateError(err) {
			return 0, schemas.ErrorResponse(400, "la categoria "+pointSaleCreate.Name+" ya existe", err)
		}
		return 0, schemas.ErrorResponse(500, "error al crear el punto de venta", err)
	}

	return pointSale.ID, nil
}

func (r *MainRepository) PointSaleUpdate(pointSaleUpdate *schemas.PointSaleUpdate) error {
	var ps models.PointSale
	if err := r.DB.First(&ps, pointSaleUpdate.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schemas.ErrorResponse(404, "punto de venta no encontrado", err)
		}
		return schemas.ErrorResponse(500, "error al obtener el punto de venta", err)
	}

	if err := r.DB.Model(&ps).Updates(models.PointSale{
		Name:        pointSaleUpdate.Name,
		Description: pointSaleUpdate.Description,
	}).Error; err != nil {
		if IsDuplicateError(err) {
			return schemas.ErrorResponse(400, "el punto de venta "+pointSaleUpdate.Name+" ya existe", err)
		}
		return schemas.ErrorResponse(500, "error al actualizar el punto de venta", err)
	}

	return nil
}

func (r *MainRepository) PointSaleDelete(id uint) error {
	err := r.DB.Where("id = ?", id).Delete(&models.PointSale{}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schemas.ErrorResponse(404, "punto de venta no encontrado", err)
		}
		return schemas.ErrorResponse(500, "error al eliminar el punto de venta", err)
	}

	return nil
}
