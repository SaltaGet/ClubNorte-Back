package models

type ReportMovement struct {
	Income             []*Income             `json:"income"`
	Expense            []*Expense            `json:"expense"`
	IncomeSportsCourts []*IncomeSportsCourts `json:"income_sports_courts"`
	ExpenseBuy         []*ExpenseBuy         `json:"expense_buy"`
	PointSale          []*PointSale          `json:"point_sale"`
}
