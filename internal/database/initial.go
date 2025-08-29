package database

import (
	"fmt"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"gorm.io/gorm"
)


func initialData(db *gorm.DB) error {
	roles := []models.Role{
		{Name: "admin"},
		{Name: "vendedor"},
		{Name: "repositor"},
	}

	for _, role := range roles {
		db.FirstOrCreate(&role, models.Role{Name: role.Name})
	}

	var rolesDB []models.Role
	db.Model(&models.Role{}).Find(&rolesDB)

	if len(rolesDB) != 3 {
		return fmt.Errorf("los roles presentan un error, se esperaban 3 y se encontraron %d", len(rolesDB))
	}

	return nil
}
