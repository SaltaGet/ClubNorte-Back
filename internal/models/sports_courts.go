package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type SportsCourt struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Code        string  `gorm:"size:50;not null;uniqueIndex"`
	Name        string  `gorm:"size:100;not null"`
	Description *string `gorm:"type:text"` 
	CreatedAt   time.Time  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
	PointSales       []PointSale `gorm:"many2many:sports_courts_point_sales;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"point_sales"`
}

func (s *SportsCourt) BeforeCreate(tx *gorm.DB) (err error) {
	s.Code = strings.ToLower(s.Code)
	return
}