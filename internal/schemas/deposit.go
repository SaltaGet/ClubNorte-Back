package schemas

type DepositResponse struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    CategoryResponse `json:"category"`
	Price       float64 `json:"price"`
	Stock       int    `json:"stock"`
}