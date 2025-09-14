package models

import "time"

type Register struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PointSaleID uint      `gorm:"not null" json:"point_sale_id"`
	PointSale   PointSale `gorm:"foreignKey:PointSaleID;references:ID" json:"point_sale"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	OpenAmount  float64   `gorm:"" json:"open_amount"`
	HourOpen    time.Time `gorm:"" json:"hour_open"`
	CloseAmount *float64   `gorm:"" json:"close_amount"`
	HourClose   *time.Time `gorm:"" json:"hour_close"`
	AmountTotal *float64   `gorm:"" json:"amount_total"`
	IsClose     bool      `gorm:"not null,default:false" json:"is_close"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
