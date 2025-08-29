package models

import "time"

type IncomeSportsCourts struct {
	ID                   uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	SportsCourtID        uint        `gorm:"not null" json:"sports_court_id"`
	SportsCourt          SportsCourt `gorm:"foreignKey:SportsCourtID;references:ID" json:"sports_court"`
	Shift                string      `gorm:"not null" json:"shift" validate:"oneof=mañana tarde noche"`
	UserID               uint        `gorm:"not null" json:"user_id"`
	User                 User        `gorm:"foreignKey:UserID;references:ID" json:"user"`
	PartialPay           *float64    `gorm:"" json:"partial_pay"`
	PartialPaymentMethod string      `gorm:"size:30;default:'efectivo'" json:"partial_payment_method" validate:"oneof=efectivo tarjeta transferencia"`
	DateTimeSign         *time.Time  `gorm:"not null" json:"date_time_sign"`
	TotalPay             float64     `gorm:"not null" json:"total_pay"`
	TotalPaymentMethod   string      `gorm:"size:30;default:'efectivo'" json:"total_payment_method" validate:"oneof=efectivo tarjeta transferencia"`
	Price                float64     `gorm:"not null" json:"price"`
	CreatedAt            time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}
