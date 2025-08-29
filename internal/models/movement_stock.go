package models

import "time"

type MovementStock struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	Amount    float64   `gorm:"not null" json:"amount"`
	FromID    uint      `gorm:"not null" json:"from_id"`
	FromType  string    `gorm:"not null" json:"from_type" validate:"oneof=deposit point_sale"`
	ToID      uint      `gorm:"not null" json:"to_id"`
	ToType    string    `gorm:"not null" json:"to_type" validate:"oneof=deposit point_sale"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
