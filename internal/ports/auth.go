package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type AuthRepository interface {
	Login(params *schemas.Login) (*models.User, error)
}

type AuthService interface {
	Login(params *schemas.Login) (string, error)
}