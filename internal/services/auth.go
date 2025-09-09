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
		return "", schemas.ErrorResponse(500, "error al generar el token", err)
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

	userContext.RoleID = user.Role.ID
	userContext.Role = user.Role.Name

	return &userContext, nil
}

func (s *AuthService) AuthPointSale(userID uint, pointSaleID uint) (*schemas.PointSaleContext, error) {
	pointSale, err := s.AuthRepository.AuthPointSale(userID, pointSaleID)
	if err != nil {
		return nil, err
	}

	var pointSaleResponse schemas.PointSaleContext
	_ = copier.Copy(&pointSaleResponse, &pointSale)

	return &pointSaleResponse, nil
}

func (s *AuthService) LoginPointSale(userID uint, pointSaleID uint) (string, error) {
	user, err := s.UserRepository.UserGetByID(userID)
	if err != nil {
		return "", err
	}

	var pointSaleResponse schemas.PointSaleResponse
	if user.Role.Name == "admin" {
		point_sale, err := s.PointSaleRepository.PointSaleGet(pointSaleID)
		if err != nil {
			return "", err
		}
		_ = copier.Copy(&pointSaleResponse, &point_sale)
	} else {
		point_sale, err := s.AuthRepository.LoginPointSale(user.ID, pointSaleID)
		if err != nil {
			return "", err
		}
		_ = copier.Copy(&pointSaleResponse, &point_sale)
	}

	var UserResponseToken schemas.UserResponseToken
	_ = copier.Copy(&UserResponseToken, &user)

	token, err := utils.GenerateToken(&UserResponseToken, &pointSaleResponse)
	if err != nil {
		return "", schemas.ErrorResponse(500, "error al generar el token", err)
	}

	return token, nil
}

func (s *AuthService) LogoutPointSale(userID uint) (string, error) {
	user, err := s.UserRepository.UserGetByID(userID)
	if err != nil {
		return "", err
	}

	var UserResponseToken schemas.UserResponseToken
	_ = copier.Copy(&UserResponseToken, &user)

	token, err := utils.GenerateToken(&UserResponseToken, nil)
	if err != nil {
		return "", schemas.ErrorResponse(500, "error al generar el token", err)
	}

	return token, nil
}

func (s *AuthService) CurrentUser(userID uint) (*schemas.UserResponse, error) {
	user, err := s.UserRepository.UserGetByID(userID)
	if err != nil {
		return nil, err
	}

	var userResponse schemas.UserResponse
	_ = copier.Copy(&userResponse, &user)

	return &userResponse, nil
}

func (s *AuthService) CurrentPointSale(ID uint) (*schemas.PointSaleResponse, error) {
	pointSale, err := s.PointSaleRepository.PointSaleGet(ID)
	if err != nil {
		return nil, err
	}

	var pointSaleResponse schemas.PointSaleResponse
	_ = copier.Copy(&pointSaleResponse, &pointSale)

	return &pointSaleResponse, nil
}