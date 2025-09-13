package services

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/jinzhu/copier"
)

func (s *SportCourtService) SportCourtGetByID(pointSaleID, id uint) (*schemas.SportCourtResponse, error) {
	sportCourt, err := s.SportCourtRepository.SportCourtGetByID(pointSaleID, id)
	if err != nil {
		return nil, err
	}

	var sportCourtResponse schemas.SportCourtResponse
	_ = copier.Copy(&sportCourtResponse, sportCourt)

	return &sportCourtResponse, nil
}

func (s *SportCourtService) SportCourtGetByCode(pointSaleID uint, code string) (*schemas.SportCourtResponse, error) {
	sportCourt, err := s.SportCourtRepository.SportCourtGetByCode(pointSaleID, code)
	if err != nil {
		return nil, err
	}

	var sportCourtResponse schemas.SportCourtResponse
	_ = copier.Copy(&sportCourtResponse, sportCourt)

	return &sportCourtResponse, nil
}

func (s *SportCourtService) SportCourtGetAllByPointSale(pointSaleID uint) ([]*schemas.SportCourtResponseDTO, error) {
	sportCourts, err := s.SportCourtRepository.SportCourtGetAllByPointSale(pointSaleID)
	if err != nil {
		return nil, err
	}

	var sportCourtsResponse []*schemas.SportCourtResponseDTO
	_ = copier.Copy(&sportCourtsResponse, sportCourts)

	return sportCourtsResponse, nil
}

func (s *SportCourtService) SportCourtGetAll(page, limit int) ([]*schemas.SportCourtResponseDTO, int64, error) {
	sportCourts, total, err := s.SportCourtRepository.SportCourtGetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var sportCourtsResponse []*schemas.SportCourtResponseDTO
	_ = copier.Copy(&sportCourtsResponse, sportCourts)

	return sportCourtsResponse, total, nil
}

func (s *SportCourtService) SportCourtCreate(pointSaleID uint, sportCourtCreate *schemas.SportCourtCreate) (uint, error) {
	return s.SportCourtRepository.SportCourtCreate(pointSaleID, sportCourtCreate)
}

func (s *SportCourtService) SportCourtUpdate(pointSaleID uint, sportCourtUpdate *schemas.SportCourtUpdate) error {
	return s.SportCourtRepository.SportCourtUpdate(pointSaleID, sportCourtUpdate)
}

func (s *SportCourtService) SportCourtDelete(pointSaleID, id uint) error {
	return s.SportCourtRepository.SportCourtDelete(pointSaleID, id)
}
