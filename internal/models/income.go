package models

import "time"

type Income struct {
	ID            uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	PointSaleID   uint         `gorm:"not null" json:"point_sale_id"`
	PointSale     PointSale    `gorm:"foreignKey:PointSaleID;references:ID" json:"point_sale"`
	UserID        uint         `gorm:"not null" json:"user_id"`
	User          User         `gorm:"foreignKey:UserID;references:ID" json:"user"`
	RegisterID    uint         `gorm:"not null" json:"register_id"`
	Register      Register     `gorm:"foreignKey:RegisterID;references:ID" json:"register"`
	Items         []IncomeItem `gorm:"foreignKey:IncomeID" json:"items"`
	Description   *string      `gorm:"size:255" json:"description"`
	Total         float64      `gorm:"not null" json:"total"`
	PaymentMethod string       `gorm:"size:30;default:'efectivo'" json:"payment_method" validate:"oneof=efectivo tarjeta transferencia"`
	CreatedAt     time.Time    `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt     time.Time    `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

type IncomeItem struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	IncomeID  uint    `gorm:"not null" json:"income_id"`
	Income    Income  `gorm:"foreignKey:IncomeID;references:ID" json:"-"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	Quantity  float64 `gorm:"not null" json:"quantity"`
	Price_Cost float64 `gorm:"not null" json:"price_cost"`
	Price     float64 `gorm:"not null" json:"price"`
	Subtotal  float64 `gorm:"not null" json:"subtotal"`
}
