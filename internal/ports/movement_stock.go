package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type MovementStockRepository interface {
	MovementStockGetByID(id uint) (*models.MovementStock, error)
	MovementStockGetAll(page, limit int) ([]*models.MovementStock, int64, error)
	MoveStock(userID uint, input *schemas.MovementStock) error
}

type MovementStockService interface {
	MovementStockGetByID(id uint) (*schemas.MovementStockResponse, error)
	MovementStockGetAll(page, limit int) ([]*schemas.MovementStockResponseDTO, int64, error)
	MoveStock(userID uint, input *schemas.MovementStock) error
}
