package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type UserRepository interface {
	UserGetByID(id uint) (*models.User, error)
	UserGetByEmail(email string) (*models.User, error)
	UserGetAll() ([]*models.User, error)
	UserCreate(userCreate *schemas.UserCreate) (uint, error)
	UserUpdate(userUpdate *schemas.UserUpdate) error
	UserDelete(id uint)	error
	UserUpdatePassword(userID uint, updatePass *schemas.UserUpdatePassword) error
}

type UserService interface {
	UserGetByID(id uint) (*schemas.UserResponse, error)
	UserGetByEmail(email string) (*schemas.UserResponse, error)
	UserGetAll() ([]*schemas.UserResponseDTO, error)
	UserCreate(userCreate *schemas.UserCreate) (uint, error)
	UserUpdate(userUpdate *schemas.UserUpdate) error
	UserDelete(id uint)	error
	UserUpdatePassword(userID uint, updatePass *schemas.UserUpdatePassword) error
}
