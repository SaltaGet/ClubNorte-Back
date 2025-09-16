package models

import "time"

type Register struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PointSaleID uint      `gorm:"not null" json:"point_sale_id"`
	PointSale   PointSale `gorm:"foreignKey:PointSaleID;references:ID" json:"point_sale"`
	UserOpenID  uint      `gorm:"not null" json:"user_open_id"`
	UserOpen    User      `gorm:"foreignKey:UserOpenID;references:ID" json:"user_open"`
	OpenAmount  float64   `gorm:"" json:"open_amount"`
	HourOpen    time.Time `gorm:"" json:"hour_open"`

	UserCloseID *uint      `json:"user_close_id"`
	UserClose   *User      `gorm:"foreignKey:UserCloseID;references:ID" json:"user_close"`
	CloseAmount *float64   `gorm:"" json:"close_amount"`
	HourClose   *time.Time `gorm:"" json:"hour_close"`

	TotalIncomeCash    *float64 `gorm:"" json:"total_income_cash"`
	TotalIncomeOthers  *float64 `gorm:"" json:"total_income_others"`
	TotalExpenseCash   *float64 `gorm:"" json:"total_expense_cash"`
	TotalExpenseOthers *float64 `gorm:"" json:"total_expense_others"`

	IsClose   bool      `gorm:"not null,default:false" json:"is_close"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}
