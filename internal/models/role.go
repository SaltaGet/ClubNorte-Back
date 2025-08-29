package models

type Role struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null;uniqueIndex;size:50" json:"name" validate:"oneof=admin vendedor repositor"`
	Users       []User `gorm:"foreignKey:RoleID"`
}