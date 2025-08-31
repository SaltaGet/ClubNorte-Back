package models

import "time"

type PointSale struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Description *string   `gorm:"size:200" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Users       []User `gorm:"many2many:user_point_sales;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"users"`
	SportsCourts       []SportsCourt `gorm:"many2many:sports_courts_point_sales;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sports_courts"`
}