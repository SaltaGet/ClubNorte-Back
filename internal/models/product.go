package models

import "time"

type Product struct {
	ID             uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	Code           string          `gorm:"size:50;not null;uniqueIndex" json:"code"`
	Name           string          `gorm:"size:100;not null" json:"name"`
	Description    *string         `gorm:"size:200" json:"description"`
	Price          float64         `gorm:"not null" json:"price"`
	CategoryID     uint            `gorm:"not null" json:"category_id"`
	Category       Category        `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
	CreatedAt      time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	StockDeposit   *StockDeposit   `gorm:"foreignKey:ProductID" json:"stock_deposit"`
	StockPointSales []*StockPointSale `gorm:"foreignKey:ProductID" json:"stock_point_sales"`
}
