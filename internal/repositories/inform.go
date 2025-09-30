package repositories

import "github.com/DanielChachagua/Club-Norte-Back/internal/models"

func (r *MainRepository) Inform() (any, error) {
  var products []*models.Product

	err := r.DB.Preload("Category").Find(&products).Error
	return products, err
}