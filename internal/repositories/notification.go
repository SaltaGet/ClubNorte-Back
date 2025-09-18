package repositories

import "github.com/DanielChachagua/Club-Norte-Back/internal/models"

func (r *MainRepository) NotificationStock() ([]*models.Product, error) {
	var products []*models.Product

	err := r.DB.
		Model(&models.Product{}).
		Joins("JOIN stock_deposits sd ON sd.product_id = products.id").
		// Where("sd.stock <= ?", 10).
		Where("sd.stock <= products.min_amount").
		Where("products.notifier = ?", true).
		Preload("StockDeposit").
		Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

