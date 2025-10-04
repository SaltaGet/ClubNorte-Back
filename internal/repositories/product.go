package repositories

import (
	"errors"

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

func (r *MainRepository) ProductGetAll(page, limit int) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64
	offset := (page - 1) * limit

	if err := r.DB.
		Model(&models.Product{}).
		Count(&total).Error; err != nil {
		return nil, 0, schemas.ErrorResponse(500, "error al contar productos", err)
	}

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
	if productCreate.Price != nil {
		product.Price = *productCreate.Price
	}
	product.CategoryID = productCreate.CategoryID
	product.Notifier = productCreate.Notifier
	product.MinAmount = productCreate.MinAmount

	if err := r.DB.Create(&product).Error; err != nil {
		if IsDuplicateError(err) {
			return 0, schemas.ErrorResponse(400, "el producto de codigo "+product.Code+" ya existe", err)
		}
		return 0, schemas.ErrorResponse(500, "error al crear el producto", err)
	}

	return product.ID, nil
}

func (r *MainRepository) ProductUpdate(product *schemas.ProductUpdate) error {
	var p models.Product
	if err := r.DB.First(&p, product.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schemas.ErrorResponse(404, "producto no encontrado", err)
		}
		return schemas.ErrorResponse(500, "error al obtener el producto", err)
	}

	if product.Price != nil {
		p.Price = *product.Price
	}

	updates := map[string]any{
		"code":        product.Code,
		"name":        product.Name,
		"description": &product.Description,
		"category_id": product.CategoryID,
		"price":       p.Price,
		"notifier":    product.Notifier,
		"min_amount":  product.MinAmount,
	}

	if err := r.DB.Model(&p).Updates(updates).Error; err != nil {
		if IsDuplicateError(err) {
			return schemas.ErrorResponse(400, "el producto de código "+product.Code+" ya existe", err)
		}
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
