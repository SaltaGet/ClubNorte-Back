package ports

import "github.com/DanielChachagua/Club-Norte-Back/internal/models"

type UserRepository interface {
	UserGetByID(id uint) (*models.User, error)
	UserGetByEmail(email string) (*models.User, error)
}

type UserService interface {
	// UserGetByID(id uint) (*models.User, error)
	// UserGetByEmail(email string) (*models.User, error)
}
