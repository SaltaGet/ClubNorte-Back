package repositories

import (
	"errors"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) SportCourtGetByID(pointSaleID, id uint) (*models.SportsCourt, error) {
	var sportCourt *models.SportsCourt

	if err := r.DB.
		Joins("JOIN sports_courts_point_sales scps ON scps.sports_court_id = sports_courts.id").
		Where("scps.point_sale_id = ?", pointSaleID).
		Preload("PointSales", "id = ?", pointSaleID).
		First(&sportCourt, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "cancha no encontrada", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener la cancha", err)
	}

	return sportCourt, nil
}

func (r *MainRepository) SportCourtGetByCode(pointSaleID uint, code string) (*models.SportsCourt, error) {
	var sportCourt *models.SportsCourt

	if err := r.DB.
		Joins("JOIN sports_courts_point_sales scps ON scps.sports_court_id = sports_courts.id").
		Where("scps.point_sale_id = ?", pointSaleID).
		Preload("PointSales", "id = ?", pointSaleID).
		Where("code = ?", code).
		First(&sportCourt).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "cancha no encontrada", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener la cancha", err)
	}

	return sportCourt, nil
}

func (r *MainRepository) SportCourtGetAllByPointSale(pointSaleID uint) ([]*models.SportsCourt, error) {
	var sportCourts []*models.SportsCourt

	if err := r.DB.
		Joins("JOIN sports_courts_point_sales scps ON scps.sports_court_id = sports_courts.id").
		Where("scps.point_sale_id = ?", pointSaleID).
		Preload("PointSales", "id = ?", pointSaleID).
		Find(&sportCourts).Error; err != nil {
		return nil, schemas.ErrorResponse(500, "error al obtener las canchas", err)
	}

	return sportCourts, nil
}

func (r *MainRepository) SportCourtGetAll(page, limit int) ([]*models.SportsCourt, int64, error) {
	var sportCourts []*models.SportsCourt
	offset := (page - 1) * limit

	if err := r.DB.
		Preload("PointSales").
		Offset(offset).
		Limit(limit).
		Find(&sportCourts).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al obtener las canchas", err)
	}

	var total int64
	if err := r.DB.Model(&models.SportsCourt{}).Count(&total).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al contar las canchas", err)
	}

	return sportCourts, total, nil
}

// func (r *MainRepository) SportsCourtsCreate(pointSaleID uint, sportCourtCreate *schemas.SportCourtCreate) (uint, error) {
// 	return r.DB.Transaction(func(tx *gorm.DB) (uint, error) {
// 		var pointSale models.PointSale
// 		if err := r.DB.Where("id = ?", pointSaleID).Find(&pointSale).Error; err != nil {
// 			return 0, schemas.ErrorResponse(404, "punto de venta no encontrado", err)
// 		}

// 		var sportCourt models.SportsCourt
// 		sportCourt.Code = sportCourtCreate.Code
// 		sportCourt.Name = sportCourtCreate.Name
// 		sportCourt.Description = sportCourtCreate.Description

// 		if err := r.DB.Create(&sportCourt).Error; err != nil {
// 			return 0, schemas.ErrorResponse(500, "error al crear la cancha", err)
// 		}

// 		return 0, nil
// 	})
// }

func (r *MainRepository) SportCourtCreate(pointSaleID uint, sportCourtCreate *schemas.SportCourtCreate) (uint, error) {
	var pointSaleIDCreate uint
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		var pointSale models.PointSale
		if err := tx.Where("id = ?", pointSaleID).First(&pointSale).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return schemas.ErrorResponse(404, "punto de venta no encontrado", err)
			}
			return schemas.ErrorResponse(500, "error al obtener el punto de venta", err)
		}

		sportCourt := models.SportsCourt{
			Code:        sportCourtCreate.Code,
			Name:        sportCourtCreate.Name,
			Description: sportCourtCreate.Description,
		}

		if err := tx.Create(&sportCourt).Error; err != nil {
			if IsDuplicateError(err) {
				return schemas.ErrorResponse(400, "la cancha "+sportCourtCreate.Name+" ya existe", err)
			}
			return schemas.ErrorResponse(500, "error al crear la cancha", err)
		}

		if err := tx.Model(&sportCourt).Association("PointSales").Append(&pointSale); err != nil {
			return schemas.ErrorResponse(500, "error al asociar la cancha al punto de venta", err)
		}

		pointSaleIDCreate = sportCourt.ID

		return nil
	})

	if err != nil {
		return 0, err
	}

	return pointSaleIDCreate, nil
}

func (r *MainRepository) SportCourtUpdate(pointSaleID uint, sportCourtUpdate *schemas.SportCourtUpdate) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var sportCourt models.SportsCourt
		if err := tx.
			Joins("JOIN sports_courts_point_sales scps ON scps.sports_court_id = sports_courts.id").
			Where("scps.point_sale_id = ?", pointSaleID).
			First(&sportCourt, sportCourtUpdate.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return schemas.ErrorResponse(404, "cancha no encontrada", err)
			}
			return schemas.ErrorResponse(500, "error al obtener la cancha", err)
		}

		sportCourt.Code = sportCourtUpdate.Code
		sportCourt.Name = sportCourtUpdate.Name
		sportCourt.Description = sportCourtUpdate.Description

		if err := tx.Save(&sportCourt).Error; err != nil {
			return schemas.ErrorResponse(500, "error al actualizar la cancha", err)
		}

		return nil
	})
}

func (r *MainRepository) SportCourtDelete(pointSaleID, courtID uint) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		// Verificar si la relación existe
		var count int64
		if err := tx.Table("sports_courts_point_sales").
			Where("point_sale_id = ? AND sports_court_id = ?", pointSaleID, courtID).
			Count(&count).Error; err != nil {
			return schemas.ErrorResponse(500, "error al verificar la relación", err)
		}

		if count == 0 {
			return schemas.ErrorResponse(404, "la cancha no está asociada a este punto de venta", nil)
		}

		// Eliminar la relación en la tabla pivote
		if err := tx.Table("sports_courts_point_sales").
			Where("point_sale_id = ? AND sports_court_id = ?", pointSaleID, courtID).
			Delete(nil).Error; err != nil {
			return schemas.ErrorResponse(500, "error al eliminar la relación", err)
		}

		// Eliminar la cancha de la tabla principal
		if err := tx.Delete(&models.SportsCourt{}, courtID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return schemas.ErrorResponse(404, "cancha no encontrada", err)
			}
			return schemas.ErrorResponse(500, "error al eliminar la cancha", err)
		}

		return nil
	})
}
