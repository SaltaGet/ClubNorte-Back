package repositories

import "github.com/DanielChachagua/Club-Norte-Back/internal/models"

func (r *MainRepository) ReportExcelGet() (any, error) {
  var products []*models.Product

	err := r.DB.Preload("Category").Find(&products).Error
	return products, err
}

func (r *MainRepository) ReportMonthGet() (any, error) {
	var income []models.Income
	if err := r.DB.Find(&income).Error; err != nil {
		return nil, err
	}
	return income, nil

}