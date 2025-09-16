package models

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	ID         uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName  string      `gorm:"not null" json:"first_name"`
	LastName   string      `gorm:"not null" json:"last_name"`
	Address    *string     `gorm:"default:null" json:"address"`
	Cellphone  *string     `gorm:"default:null" json:"cellphone"`
	Email      string      `gorm:"not null;uniqueIndex;size:100" json:"email"`
	Username   string      `gorm:"not null;uniqueIndex;size:50" json:"username"`
	Password   string      `gorm:"not null" json:"password"`
	IsActive   bool        `gorm:"default:true" json:"is_active"`
	IsAdmin    bool        `gorm:"default:false" json:"is_admin"`
	RoleID     uint        `gorm:"not null" json:"role_id"`
	Role       Role        `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	CreatedAt   time.Time  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
	PointSales []PointSale `gorm:"many2many:user_point_sales;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"point_sales"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
    if tx.Statement.Changed("Password") {
        hashedPassword, err := utils.HashPassword(u.Password)
        if err != nil {
            return err
        }
        u.Password = hashedPassword
    }
    return
}

