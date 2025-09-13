package repositories

import (
	"errors"
	"fmt"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) ProductGetByID(id uint) (*models.Product, error) {
	var product *models.Product

	if err := r.DB.Preload("Category").
		Preload("StockPointSales").
		Preload("StockPointSales.PointSale").
		Preload("StockDeposit").
		First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "producto no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener el producto", err)
	}

	return product, nil
}

func (r *MainRepository) ProductGetByCode(code string) (*models.Product, error) {
	var product *models.Product

	if err := r.DB.
		Preload("Category").
		Preload("StockPointSales").
		Preload("StockPointSales.PointSale").
		Preload("StockDeposit").
		Where("code = ?", code).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "producto no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener el producto", err)
	}

	return product, nil
}

func (r *MainRepository) ProductGetByCategoryID(categoryID uint) ([]*models.Product, error) {
	var products []*models.Product

	if err := r.DB.
		Preload("Category").
		Preload("StockPointSales").
		Preload("StockPointSales.PointSale").
		Preload("StockDeposit").
		Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		return nil, schemas.ErrorResponse(500, "error al obtener productos", err)
	}

	return products, nil
}

func (r *MainRepository) ProductGetByName(name string) ([]*models.Product, error) {
	var products []*models.Product

	if err := r.DB.
		Preload("Category").
		Preload("StockPointSales").
		Preload("StockPointSales.PointSale").
		Preload("StockDeposit").
		Where("name LIKE ?", "%"+name+"%").Find(&products).Error; err != nil {
		return nil, schemas.ErrorResponse(500, "error al obtener productos", err)
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

func (r *MainRepository) ProductGetAll(page, limit int) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64
	offset := (page - 1) * limit

	// Contar los productos disponibles en el punto de venta
	if err := r.DB.
		Model(&models.Product{}).
		Count(&total).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al contar productos", err)
	}

	// Obtener productos con la categoría y el stock específico del punto de venta
	if err := r.DB.
		Preload("Category").
		Preload("StockPointSales").
		Preload("StockPointSales.PointSale").
		Preload("StockDeposit").
		Offset(offset).
		Limit(limit).
		Find(&products).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al obtener productos", err)
	}

	return products, total, nil
}

func (r *MainRepository) ProductCreate(productCreate *schemas.ProductCreate) (uint, error) {
	var product models.Product

	product.Name = productCreate.Name
	product.Code = productCreate.Code
	product.Description = productCreate.Description
	product.Price = productCreate.Price
	product.CategoryID = productCreate.CategoryID

	if err := r.DB.Create(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return 0, schemas.ErrorResponse(400, "el producto de codigo "+product.Code+" ya existe", err)
		}
		return 0, schemas.ErrorResponse(500, "error al crear el producto", err)
	}

	return product.ID, nil
}

func (r *MainRepository) ProductUpdate(product *schemas.ProductUpdate) error {
	var exists bool

	if err := r.DB.Model(&models.Product{}).
		Select("count(*) > 0").
		Where("id = ?", product.ID).
		Find(&exists).Error; err != nil {
		return schemas.ErrorResponse(500, "error al obtener el producto", err)
	}

	if !exists {
		return schemas.ErrorResponse(404, "producto no encontrado", fmt.Errorf("producto no encontrado id: %d", product.ID))
	}

	if err := r.DB.Model(&models.Product{}).
		Where("id = ?", product.ID).
		Updates(product).Error; err != nil {
		return schemas.ErrorResponse(500, "error al actualizar el producto", err)
	}

	return nil
}

func (r *MainRepository) ProductDelete(id uint) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("product_id = ?", id).Delete(&models.StockPointSale{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return schemas.ErrorResponse(404, "producto no encontrado", err)
			}
			return schemas.ErrorResponse(500, "error al eliminar el producto", err)
		}

		if err := tx.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return schemas.ErrorResponse(404, "producto no encontrado", err)
			}
			return schemas.ErrorResponse(500, "error al eliminar el producto", err)
		}
		return nil
	})
}
