package models

import "time"

type Register struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PointSaleID uint      `gorm:"not null" json:"point_sale_id"`
	PointSale   PointSale `gorm:"foreignKey:PointSaleID;references:ID" json:"point_sale"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	OpenAmount  float64   `gorm:"not null" json:"open_amount"`
	HourOpen    time.Time `gorm:"not null" json:"hour_open"`
	CloseAmount float64   `gorm:"not null" json:"close_amount"`
	HourClose   time.Time `gorm:"not null" json:"hour_close"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
