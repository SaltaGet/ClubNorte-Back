package models

import "time"

type StockDeposite struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	Quantity  float64   `gorm:"not null;default:0" json:"quantity"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type StockPointSale struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID   uint      `gorm:"not null" json:"product_id"`
	Product     Product   `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	PointSaleID uint      `gorm:"not null" json:"point_sale_id"`
	PointSale   PointSale `gorm:"foreignKey:PointSaleID;references:ID" json:"point_sale"`
	Quantity    float64   `gorm:"not null;default:0" json:"quantity"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
