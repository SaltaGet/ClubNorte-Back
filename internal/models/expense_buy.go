package models

import "time"

type ExpenseBuy struct {
	ID              uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          uint              `gorm:"not null" json:"user_id"`
	User            User              `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Description     *string           `gorm:"size:255" json:"description"`
	ItemExpenseBuys []*ItemExpenseBuy `gorm:"foreignKey:ExpenseBuyID" json:"item_expense_buys"`
	PaymentMethod   string            `gorm:"size:30;default:'efectivo'" json:"payment_method" validate:"oneof=efectivo tarjeta transferencia"`
	Total           float64           `gorm:"total" json:"total"`
	CreatedAt       time.Time         `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt       time.Time         `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

type ItemExpenseBuy struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	ExpenseBuyID uint       `gorm:"not null" json:"expense_buy_id"`
	ExpenseBuy   ExpenseBuy `gorm:"foreignKey:ExpenseBuyID;references:ID" json:"expense_buy"`
	ProductID    uint       `gorm:"not null" json:"product_id"`
	Product      Product    `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	Quantity     float64       `gorm:"not null" json:"quantity"`
	Price        float64    `gorm:"not null" json:"price"`
	Subtotal     float64    `gorm:"not null" json:"subtotal"`
	CreatedAt    time.Time  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime:milli" json:"updated_at"`
}
