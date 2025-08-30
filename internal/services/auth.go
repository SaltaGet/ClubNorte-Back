package services

import (
	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/DanielChachagua/Club-Norte-Back/internal/utils"
	"github.com/jinzhu/copier"	
)

func (s *AuthService) Login(params *schemas.Login) (string, error) {
	user, err := s.AuthRepository.Login(params)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(params.Password, user.Password) {
		return "", schemas.ErrorResponse(401, "Credenciales incorrectas", nil)
	}

	var userResponse schemas.UserResponseToken
	_ = copier.Copy(&userResponse, &user)

	token, err := utils.GenerateToken(&userResponse, nil)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) AuthUser(email string) (*schemas.UserContext, error) {
	user, err := s.AuthRepository.AuthUser(email)
	if err != nil {
		return nil, err
	}

	var userContext schemas.UserContext
	_ = copier.Copy(&userContext, &user)


	return &userContext, nil
}

func (s *AuthService) AuthPointSale(userID uint, pointSaleID uint) (*schemas.PointSaleResponse, error) {
	pointSale, err := s.AuthRepository.AuthPointSale(userID, pointSaleID)
	if err != nil {
		return nil, err
	}

	var pointSaleResponse schemas.PointSaleResponse
	_ = copier.Copy(&pointSaleResponse, &pointSale)

	return &pointSaleResponse, nil
}