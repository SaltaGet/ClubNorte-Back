package repositories

import (
	"errors"
	"fmt"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) CategoryGetByID(id uint) (*models.Category, error) {
	var category *models.Category

	if err := r.DB.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("categoria no encontrada")
		}
		return nil, err
	}

	return category, nil
}

func (r *MainRepository) CategoryGetAll() ([]*models.Category, error) {
	var categories []*models.Category

	if err := r.DB.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *MainRepository) CategoryCreate(categoryCreate *schemas.CategoryCreate) (uint, error) {
	var category models.Category

	category.Name = categoryCreate.Name

	if err := r.DB.Create(&category).Error; err != nil {
		return 0, err
	}

	return category.ID, nil
}

func (r *MainRepository) CategoryUpdate(categoryUpdate *schemas.CategoryUpdate) error {
	var exists bool

	if err := r.DB.Model(&models.Category{}).
		Select("count(*) > 0").
		Where("id = ?", categoryUpdate.ID).
		Find(&exists).Error; err != nil {
		return err
	}

	if !exists {
		return gorm.ErrRecordNotFound
	}

	return r.DB.Model(&models.Category{}).
		Where("id = ?", categoryUpdate.ID).
		Updates(map[string]interface{}{
			"name": categoryUpdate.Name,
		}).Error
}

func (r *MainRepository) CategoryDelete(id uint) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.Category{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("categoria no encontrada")
		}
		return err
	}

	return nil
}
