package ports

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type RegisterRepository interface {
	RegisterOpen(pointSaleID uint, userID uint, amountOpen schemas.RegisterOpen) error
	RegisterClose(pointSaleID uint, userID uint, amountOpen schemas.RegisterClose) error
	RegisterInform(pointSaleID uint, userID uint, fromDate, toDate time.Time) ([]*models.Register, error)
	RegisterExistOpen(pointSaleID uint) (bool, error)
}

type RegisterService interface {
	RegisterOpen(pointSaleID uint, userID uint, amountOpen schemas.RegisterOpen) error
	RegisterClose(pointSaleID uint, userID uint, amountOpen schemas.RegisterClose) error
	RegisterInform(pointSaleID uint, userID uint, fromDate, toDate time.Time) ([]*schemas.RegisterInformResponse, error)
	RegisterExistOpen(pointSaleID uint) (bool, error)
}
