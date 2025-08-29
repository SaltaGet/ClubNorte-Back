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

	var userResponse schemas.UserResponseToken
	_ = copier.Copy(&userResponse, &user)

	token, err := utils.GenerateToken(&userResponse, nil)
	if err != nil {
		return "", err
	}

	return token, nil
}
