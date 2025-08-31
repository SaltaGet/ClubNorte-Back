package services

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/jinzhu/copier"
)

func (s *CategoryService) CategoryGetByID(id uint) (*schemas.CategoryResponse, error) {
	category, err := s.CategoryRepository.CategoryGetByID(id)
	if err != nil {
		return nil, err
	}

	var categoryResponse schemas.CategoryResponse
	_ = copier.Copy(&categoryResponse, &category)

	return &categoryResponse, nil
}

func (s *CategoryService) CategoryGetAll() ([]*schemas.CategoryResponse, error) {
	categories, err := s.CategoryRepository.CategoryGetAll()
	if err != nil {
		return nil, err
	}

	var categoriesResponse []*schemas.CategoryResponse
	_ = copier.Copy(&categoriesResponse, &categories)

	return categoriesResponse, nil
}

func (s *CategoryService) CategoryCreate(categoryCreate *schemas.CategoryCreate) (uint, error) {
	return s.CategoryRepository.CategoryCreate(categoryCreate)
}

func (s *CategoryService) CategoryUpdate(categoryUpdate *schemas.CategoryUpdate) error {
	return s.CategoryRepository.CategoryUpdate(categoryUpdate)
}

func (s *CategoryService) CategoryDelete(id uint) error {
	return s.CategoryRepository.CategoryDelete(id)
}