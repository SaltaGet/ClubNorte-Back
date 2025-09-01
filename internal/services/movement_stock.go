package services

func (m *MovementStockService) MovementStockDepositToPointSale(userID, pointSaleID, productID uint, amount float64) error {
	return m.MovementStockRepository.MovementStockDepositToPointSale(userID, pointSaleID, productID, amount)
}

func (m *MovementStockService) MovementStockPointSaleToPointSale(userID, fromPointSaleID, toPointSaleID, productID uint, amount float64) error {
	return m.MovementStockRepository.MovementStockPointSaleToPointSale(userID, fromPointSaleID, toPointSaleID, productID, amount)
}

func (m *MovementStockService) MovementStockPointSaleToDeposit(userID, pointSaleID, productID uint, amount float64) error {
	return m.MovementStockRepository.MovementStockPointSaleToDeposit(userID, pointSaleID, productID, amount)
}