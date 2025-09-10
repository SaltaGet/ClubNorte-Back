package schemas

type PointSaleStock struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Stock float64 `json:"stock"`
}

type StockDepositResponse struct {
	ID          uint   `json:"id"`
	Stock       float64    `json:"stock"`
}