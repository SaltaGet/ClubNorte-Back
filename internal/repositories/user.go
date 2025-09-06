package repositories

import (
	"errors"
	"fmt"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) UserGetByID(id uint) (*models.User, error) {
	var user *models.User

	if err := r.DB.Preload("Role").Preload("PointSales").Where("id = ? AND is_admin = ?", id, false).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("punto de venta no encontrado")
		}
		return nil, err
	}

	return user, nil
}

func (r *MainRepository) UserGetByEmail(email string) (*models.User, error) {
	var user *models.User

	if err := r.DB.Preload("Role").Preload("PointSales").Where("email = ? AND is_admin = ?", email, false).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("punto de venta no encontrado")
		}
		return nil, err
	}

	return user, nil
}

func (r *MainRepository) UserGetAll() ([]*models.User, error) {
	var users []*models.User

	if err := r.DB.Preload("Role").Where("is_admin = ?", false).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *MainRepository) UserCreate(userCreate *schemas.UserCreate) (uint, error) {
	var pointSales []models.PointSale

	if err := r.DB.Where("id IN (?)", userCreate.PointSaleIDs).Find(&pointSales).Error; err != nil {
		return 0, err
	}

	user := &models.User{
		FirstName:  userCreate.FirstName,
		LastName:   userCreate.LastName,
		Address:    userCreate.Address,
		Cellphone:  userCreate.Cellphone,
		Username:   userCreate.Username,
		Email:      userCreate.Email,
		Password:   userCreate.Password,
		RoleID:     userCreate.RoleID,
		PointSales: pointSales,
	}

	if err := r.DB.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r *MainRepository) UserUpdate(userUpdate *schemas.UserUpdate) error {
    return r.DB.Transaction(func(tx *gorm.DB) error {
        var pointSales []models.PointSale
        if err := tx.Where("id IN (?)", userUpdate.PointSaleIDs).Find(&pointSales).Error; err != nil {
            return err
        }

        // Actualizar datos básicos
        if err := tx.Model(&models.User{}).
            Where("id = ?", userUpdate.ID).
            Updates(models.User{
                FirstName: userUpdate.FirstName,
                LastName:  userUpdate.LastName,
                Address:   userUpdate.Address,
                Cellphone: userUpdate.Cellphone,
                Username:  userUpdate.Username,
                Email:     userUpdate.Email,
                RoleID:    userUpdate.RoleID,
            }).Error; err != nil {
            return err
        }

        // Actualizar relación con PointSales
        if err := tx.Model(&models.User{ID: userUpdate.ID}).
            Association("PointSales").
            Replace(pointSales); err != nil {
            return err
        }

        return nil
    })
}


func (r *MainRepository) UserDelete(id uint) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("usuario no encontrado")
		}
		return err
	}
	return nil
}

func (r *MainRepository) UserUpdatePassword(userID uint, userUpdatePassword *schemas.UserUpdatePassword) error {
	var user models.User
	if err := r.DB.First(&user, userID).Error; err != nil {
		return err
	}

	user.Password = userUpdatePassword.NewPassword
	return r.DB.Save(&user).Error
}
