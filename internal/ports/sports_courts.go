package ports

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/models"
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
)

type SportCourtRepository interface {
	SportCourtGetByID(pointSaleID, id uint) (*models.SportsCourt, error)
	SportCourtGetByCode(pointSaleID uint, code string) (*models.SportsCourt, error)
	SportCourtGetAll(page, limit int) ([]*models.SportsCourt, int64, error)
	SportCourtGetAllByPointSale(pointSaleID uint) ([]*models.SportsCourt, error)	
	SportCourtCreate(pointSaleID uint, sportCourtCreate *schemas.SportCourtCreate) (uint, error)
	SportCourtUpdate(pointSaleID uint, sportCourtUpdate *schemas.SportCourtUpdate) error
	SportCourtDelete(pointSaleID uint, id uint) error
}

type SportCourtService interface {
	SportCourtGetByID(pointSaleID, id uint) (*schemas.SportCourtResponse, error)
	SportCourtGetByCode(pointSaleID uint, code string) (*schemas.SportCourtResponse, error)
	SportCourtGetAll(page, limit int) ([]*schemas.SportCourtResponseDTO, int64, error)
	SportCourtGetAllByPointSale(pointSaleID uint) ([]*schemas.SportCourtResponseDTO, error)	
	SportCourtCreate(pointSaleID uint, sportCourtCreate *schemas.SportCourtCreate) (uint, error)
	SportCourtUpdate(pointSaleID uint, sportCourtUpdate *schemas.SportCourtUpdate) error
	SportCourtDelete(pointSaleID uint, id uint) error
}