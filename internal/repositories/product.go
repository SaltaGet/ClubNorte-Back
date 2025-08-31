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

	if err := r.DB.Preload("Category").First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("producto no encontrado")
		}
		return nil, err
	}

	return product, nil
}

func (r *MainRepository) ProductGetByCode(code string) (*models.Product, error) {
	var product *models.Product

	if err := r.DB.Preload("Category").Where("code = ?", code).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("producto no encontrado")
		}
		return nil, err
	}

	return product, nil
}

func (r *MainRepository) ProductGetByCategoryID(categoryID uint) ([]*models.Product, error) {
	var products []*models.Product

	if err := r.DB.Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *MainRepository) ProductGetByName(name string) ([]*models.Product, error) {
	var products []*models.Product

	if err := r.DB.Preload("Category").Where("name LIKE ?", "%"+name+"%").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *MainRepository) ProductGetAll(page, limit int) ([]*models.Product, int64, error) {
	var products []*models.Product
	var total int64

	// Contar el total de registros (para saber cuántas páginas hay)
	if err := r.DB.Model(&models.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Calcular offset
	offset := (page - 1) * limit

	// Obtener productos con paginación
	if err := r.DB.
		Preload("Category").
		Limit(limit).
		Offset(offset).
		Find(&products).Error; err != nil {
		return nil, 0, err
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
		return 0, err
	}

	return product.ID, nil
}

func (r *MainRepository) ProductUpdate(product *schemas.ProductUpdate) error {
	var exists bool

	if err := r.DB.Model(&models.Product{}).
		Select("count(*) > 0").
		Where("id = ?", product.ID).
		Find(&exists).Error; err != nil {
		return err
	}

	if !exists {
		return gorm.ErrRecordNotFound
	}

	return r.DB.Model(&models.Product{}).
		Where("id = ?", product.ID).
		Updates(product).Error
}

func (r *MainRepository) ProductDelete(id uint) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("producto no encontrado")
		}
		return err
	}

	return nil
}
