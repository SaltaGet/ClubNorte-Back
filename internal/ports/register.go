package ports

import (
	"time"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type RegisterRepository interface {
	RegisterOpen(pointSaleID uint, userID uint, amountOpen schemas.RegisterOpen) error
	RegisterClose(pointSaleID uint, userID uint, amountOpen schemas.RegisterClose) error
	RegisterInform(pointSaleID uint, userID uint, dateInform time.Time) (*schemas.RegisterInform, error)
}

type RegisterService interface {
	RegisterOpen(pointSaleID uint, userID uint, amountOpen schemas.RegisterOpen) error
	RegisterClose(pointSaleID uint, userID uint, amountOpen schemas.RegisterClose) error
	RegisterInform(pointSaleID uint, userID uint, dateInform time.Time) (*schemas.RegisterInform, error)
}