package services

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/jinzhu/copier"
)

func (m *MovementStockService) MovementStockGetByID(id uint) (*schemas.MovementStockResponse, error) {
	movement, err := m.MovementStockRepository.MovementStockGetByID(id)
	if err != nil {
		return nil, err
	}

	var response schemas.MovementStockResponse
	_ = copier.Copy(&response, movement)

	return &response, nil
}

func (m *MovementStockService) MovementStockGetAll(page, limit int) ([]*schemas.MovementStockResponseDTO, int64, error) {
	movements, total, err := m.MovementStockRepository.MovementStockGetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var movementsResponse []*schemas.MovementStockResponseDTO
	_ = copier.Copy(&movementsResponse, &movements)

	return movementsResponse, total, nil
}

func (m *MovementStockService) MoveStock(userID uint, input *schemas.MovementStock) error {
	return m.MovementStockRepository.MoveStock(userID, input)
}