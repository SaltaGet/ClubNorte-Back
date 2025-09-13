package repositories

import (
	"errors"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) StockProductGetByID(pointSaleID, id uint) (*models.Product, error) {
	var product *models.Product

	if err := r.DB.
		Joins("JOIN stock_point_sales sps ON sps.product_id = products.id").
		Where("sps.point_sale_id = ?", pointSaleID).
		Preload("Category").
		Preload("StockPointSales", "point_sale_id = ?", pointSaleID).
		Preload("StockPointSales.PointSale").
		First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "producto no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener producto", err)
	}

	return product, nil
}

func (r *MainRepository) StockProductGetByCode(pointSaleID uint, code string) (*models.Product, error) {
	var product *models.Product

	// if err := r.DB.Preload("Category").Preload("StockPointSale").Where("code = ?", code).First(&product).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return nil, schemas.ErrorResponse(404, "producto no encontrado", err)
	// 	}
	// 	return nil, schemas.ErrorResponse(500, "error al obtener el producto", err)
	// }
	if err := r.DB.
		Joins("JOIN stock_point_sales sps ON sps.product_id = products.id").
		Where("sps.point_sale_id = ?", pointSaleID).
		Preload("Category").
		Preload("StockPointSales", "point_sale_id = ?", pointSaleID).
		Preload("StockPointSales.PointSale").
		Where("code = ?", code).
		First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "producto no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al contar productos", err)
	}

	return product, nil
}

func (r *MainRepository) StockProductGetByCategoryID(pointSaleID, categoryID uint) ([]*models.Product, error) {
	var products []*models.Product

	// if err := r.DB.Preload("Category").Preload("StockPointSale").Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
	// 	return nil, schemas.ErrorResponse(500, "error al obtener productos", err)
	// }
	if err := r.DB.
		Joins("JOIN stock_point_sales sps ON sps.product_id = products.id").
		Where("sps.point_sale_id = ?", pointSaleID).
		Preload("Category").
		Preload("StockPointSales", "point_sale_id = ?", pointSaleID).
		Preload("StockPointSales.PointSale").
		Where("category_id = ?", categoryID).
		Find(&products).Error; err != nil {
		return nil, schemas.ErrorResponse(500, "error al contar productos", err)
	}

	return products, nil
}

func (r *MainRepository) StockProductGetByName(pointSaleID uint, name string) ([]*models.Product, error) {
	var products []*models.Product

	// if err := r.DB.Preload("Category").Preload("StockPointSale").Where("name LIKE ?", "%"+name+"%").Find(&products).Error; err != nil {
	// 	return nil, schemas.ErrorResponse(500, "error al obtener productos", err)
	// }
	if err := r.DB.
		Joins("JOIN stock_point_sales sps ON sps.product_id = products.id").
		Where("sps.point_sale_id = ?", pointSaleID).
		Preload("Category").
		Preload("StockPointSales", "point_sale_id = ?", pointSaleID).
		Preload("StockPointSales.PointSale").
		Where("name LIKE ?", "%"+name+"%").
		Find(&products).Error; err != nil {
		return nil, schemas.ErrorResponse(500, "error al contar productos", err)
	}

	return products, nil
}

// func (r *MainRepository) ProductGetAll(pointSaleID uint, page, limit int) ([]*models.Product, int64, error) {
// 	var products []*models.Product
// 	var total int64
// 	offset := (page - 1) * limit

// 	subQuery := r.DB.
// 		Table("stock_point_sales").
// 		Select("product_id").
// 		Where("point_sale_id = ?", pointSaleID)
// 	// if err := r.DB.Preload("StockPointSale").Where("stock_point_sale = ?", pointSaleID).Model(&models.Product{}).Count(&total).Error; err != nil {
// 	// 	return nil, 0, err
// 	// }

// 	if err := r.DB.
// 		Model(&models.Product{}).
// 		Where("id IN (?)", subQuery).
// 		Count(&total).Error; err != nil {
// 		return nil, 0, err
// 	}

// 	// if err := r.DB.
// 	// 	Preload("Category").
// 	// 	Preload("StockPointSale").
// 	// 	Limit(limit).
// 	// 	Offset(offset).
// 	// 	Find(&products).Error; err != nil {
// 	// 	return nil, 0, err
// 	// }
// 	if err := r.DB.
// 		Preload("Category").
// 		Preload("StockPointSale", "point_sale_id = ?", pointSaleID).
// 		Where("id IN (?)", subQuery).
// 		Limit(limit).
// 		Offset(offset).
// 		Find(&products).Error; err != nil {
// 		return nil, 0, err
// 	}

// 	return products, total, nil
// }

func (r *MainRepository) StockProductGetAll(pointSaleID uint, page, limit int) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64
	offset := (page - 1) * limit

	// Contar los productos disponibles en el punto de venta
	// if err := r.DB.
	// 	Model(&models.Product{}).
	// 	Preload("StockPointSales", "point_sale_id = ?", pointSaleID).
	// 	Count(&total).Error; err != nil {
	// 	return nil, 0, schemas.ErrorResponse(500, "error al contar productos", err)
	// }
	if err := r.DB.
		Model(&models.Product{}).
		Joins("JOIN stock_point_sales sps ON sps.product_id = products.id").
		Where("sps.point_sale_id = ?", pointSaleID).
		Preload("Category").
		Count(&total).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al obtener productos", err)
	}

	// Obtener productos con la categoría y el stock específico del punto de venta
	// if err := r.DB.
	// 	Preload("Category").
	// 	Preload("StockPointSales", "point_sale_id = ?", pointSaleID).
	// 	Preload("StockPointSales.PointSale").
	// 	Offset(offset).
	// 	Limit(limit).
	// 	Find(&products).Error; err != nil {
	// 	return nil, 0, schemas.ErrorResponse(500, "error al obtener productos", err)
	// }
	if err := r.DB.
		Joins("JOIN stock_point_sales sps ON sps.product_id = products.id").
		Where("sps.point_sale_id = ?", pointSaleID).
		Preload("Category").
		Preload("StockPointSales", "point_sale_id = ?", pointSaleID).
		Preload("StockPointSales.PointSale").
		Offset(offset).
		Limit(limit).
		Find(&products).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al obtener productos", err)
	}

	return products, total, nil
}
