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
		if IsDuplicateError(err) {
			field := DuplicateField(err)
			switch field {
			case "email":
				return 0, schemas.ErrorResponse(400, "el email "+user.Email+" ya existe", err)
			case "username":
				return 0, schemas.ErrorResponse(400, "el username "+user.Username+" ya existe", err)
			default:
				return 0, schemas.ErrorResponse(400, "ya existe un registro con ese valor único", err)
			}
		}
		return 0, schemas.ErrorResponse(500, "error al crear el usuario", err)
	}

	return user.ID, nil
}

func (r *MainRepository) UserUpdate(userUpdate *schemas.UserUpdate) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var u models.User
		if err := tx.First(&u, userUpdate.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return schemas.ErrorResponse(404, "usuario no encontrado", err)
			}
			return schemas.ErrorResponse(500, "error al obtener el usuario", err)
		}

		var pointSales []models.PointSale
		if err := tx.Where("id IN (?)", userUpdate.PointSaleIDs).Find(&pointSales).Error; err != nil {
			return err
		}

		updates := map[string]any{
			"first_name": userUpdate.FirstName,
			"last_name":  userUpdate.LastName,
			"address":    userUpdate.Address,
			"cellphone":  userUpdate.Cellphone,
			"username":   userUpdate.Username,
			"email":      userUpdate.Email,
			"role_id":    userUpdate.RoleID,
			"is_active":  userUpdate.IsActive,
		}

		if err := tx.Model(&u).
			Where("id = ?", userUpdate.ID).
			Updates(updates).Error; err != nil {
			if IsDuplicateError(err) {
				field := DuplicateField(err)
				switch field {
				case "email":
					return schemas.ErrorResponse(400, "el email "+userUpdate.Email+" ya existe", err)
				case "username":
					return schemas.ErrorResponse(400, "el username "+userUpdate.Username+" ya existe", err)
				default:
					return schemas.ErrorResponse(400, "ya existe un registro con ese valor único", err)
				}
			}
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

func (r *MainRepository) UserUpdateIsActive(userID uint) error {
	var user models.User
	if err := r.DB.First(&user, userID).Error; err != nil {
		return schemas.ErrorResponse(404, "usuario no encontrado", err)
	}

	user.IsActive = !user.IsActive
	if err := r.DB.Save(&user).Error; err != nil {
		return schemas.ErrorResponse(500, "error al actualizar el estado del usuario", err)
	}

	return nil
}
