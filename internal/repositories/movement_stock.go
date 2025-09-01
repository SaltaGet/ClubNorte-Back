package repositories

func (r *MainRepository) MovementStockDepositToPointSale(userID, pointSaleID, productID uint, amount float64) error {
	return nil
}

func (r *MainRepository) MovementStockPointSaleToPointSale(userID, fromPointSaleID, toPointSaleID, productID uint, amount float64) error {
	return nil
}

func (r *MainRepository) MovementStockPointSaleToDeposit(userID, pointSaleID, productID uint, amount float64) error {
	return nil
}
