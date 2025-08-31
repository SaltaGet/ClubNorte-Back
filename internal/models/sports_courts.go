package models

import "time"

type SportsCourt struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Code        string  `gorm:"size:50;not null;uniqueIndex"`
	Name        string  `gorm:"size:100;not null"`
	Description *string `gorm:"type:text"` 
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	PointSales       []User `gorm:"many2many:sports_courts_point_sales;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"point_sales"`
}