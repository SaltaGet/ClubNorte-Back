package models

import "time"

type Income struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PointSaleID   uint    `gorm:"not null" json:"point_sale_id"`
	PointSale     PointSale `gorm:"foreignKey:PointSaleID;references:ID" json:"point_sale"`
	UserID        uint    `gorm:"not null" json:"user_id"`
	User          User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	RegisterID    uint    `gorm:"not null" json:"register_id"`
	Register      Register  `gorm:"foreignKey:RegisterID;references:ID" json:"register"`
	ProductID     uint    `gorm:"not null" json:"product_id"`
	Product       Product  `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	Description   *string    `gorm:"size:255" json:"description"`
	Amount        float64   `gorm:"not null" json:"amount"`
	PaymentMethod string    `gorm:"size:30;default:'efectivo'" json:"payment_method" validate:"oneof=efectivo tarjeta transferencia"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
