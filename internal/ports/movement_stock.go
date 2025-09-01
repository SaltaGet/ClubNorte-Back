package ports

type MovementStockRepository interface {
	MovementStockDepositToPointSale(userID, pointSaleID, productID uint, amount float64) error
	MovementStockPointSaleToPointSale(userID, fromPointSaleID, toPointSaleID, productID uint, amount float64) error
	MovementStockPointSaleToDeposit(userID, pointSaleID, productID uint, amount float64) error
}

type MovementStockService interface {
	MovementStockDepositToPointSale(userID, pointSaleID, productID uint, amount float64) error
	MovementStockPointSaleToPointSale(userID, fromPointSaleID, toPointSaleID, productID uint, amount float64) error
	MovementStockPointSaleToDeposit(userID, pointSaleID, productID uint, amount float64) error
}