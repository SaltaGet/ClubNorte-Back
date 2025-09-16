package models

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"size:100;not null" json:"name"`
	CreatedAt   time.Time  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}