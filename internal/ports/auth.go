package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type AuthRepository interface {
	Login(params *schemas.Login) (*models.User, error)
	AuthUser(email string) (*models.User, error)
	AuthPointSale(userID uint, pointSaleID uint) (*models.PointSale, error)
}

type AuthService interface {
	Login(params *schemas.Login) (string, error)
	AuthUser(email string) (*schemas.UserContext, error)
	AuthPointSale(userID uint, pointSaleID uint) (*schemas.PointSaleResponse, error)
}