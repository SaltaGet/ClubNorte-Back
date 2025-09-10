package repositories

import (
	"errors"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"gorm.io/gorm"
)

func (r *MainRepository) UserGetByID(id uint) (*models.User, error) {
	var user *models.User

	if err := r.DB.Preload("Role").Preload("PointSales").Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "usuario no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener el usuario", err)
	}

	return user, nil
}

func (r *MainRepository) UserGetByEmail(email string) (*models.User, error) {
	var user *models.User

	if err := r.DB.Preload("Role").Preload("PointSales").Where("email = ? AND is_admin = ?", email, false).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, schemas.ErrorResponse(404, "usuario no encontrado", err)
		}
		return nil, schemas.ErrorResponse(500, "error al obtener el usuario", err)
	}

	return user, nil
}

func (r *MainRepository) UserGetAll() ([]*models.User, error) {
	var users []*models.User

	if err := r.DB.Preload("Role").Where("is_admin = ?", false).Find(&users).Error; err != nil {
		return nil, schemas.ErrorResponse(500, "error al obtener los usuarios", err)
	}

	return users, nil
}

func (r *MainRepository) UserCreate(userCreate *schemas.UserCreate) (uint, error) {
	var pointSales []models.PointSale

	if err := r.DB.Where("id IN (?)", userCreate.PointSaleIDs).Find(&pointSales).Error; err != nil {
		return 0, schemas.ErrorResponse(500, "error al obtener los puntos de venta", err)
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
		return 0, schemas.ErrorResponse(500, "error al crear el usuario", err)
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
            return schemas.ErrorResponse(500, "error al actualizar el usuario", err)
        }

        // Actualizar relación con PointSales
        if err := tx.Model(&models.User{ID: userUpdate.ID}).
            Association("PointSales").
            Replace(pointSales); err != nil {
            return schemas.ErrorResponse(500, "error al actualizar la relación con los puntos de venta", err)
        }

        return nil
    })
}


func (r *MainRepository) UserDelete(id uint) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schemas.ErrorResponse(404, "usuario no encontrado", err)
		}
		return schemas.ErrorResponse(500, "error al eliminar el usuario", err)
	}
	return nil
}

func (r *MainRepository) UserUpdatePassword(userID uint, userUpdatePassword *schemas.UserUpdatePassword) error {
	var user models.User
	if err := r.DB.First(&user, userID).Error; err != nil {
		return schemas.ErrorResponse(404, "usuario no encontrado", err)
	}

	user.Password = userUpdatePassword.NewPassword
	if err := r.DB.Save(&user).Error; err != nil {
		return schemas.ErrorResponse(500, "error al actualizar la contraseña", err)
	}

	return nil
}
