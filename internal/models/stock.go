package models

import "time"

type StockDeposit struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Stock  float64   `gorm:"not null;default:0" json:"stock"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type StockPointSale struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID   uint      `gorm:"not null" json:"product_id"`
	// Product     Product   `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	PointSaleID uint      `gorm:"not null" json:"point_sale_id"`
	PointSale   PointSale `gorm:"foreignKey:PointSaleID;references:ID" json:"point_sale"`
	Stock    float64   `gorm:"not null;default:0" json:"stock"`
	Ignored     bool      `gorm:"not null;default:false" json:"ignored"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
