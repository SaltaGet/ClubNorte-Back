package services

import (
	"fmt"

	"github.com/DanielChachagua/Club-Norte-Back/internal/schemas"
	"github.com/DanielChachagua/Club-Norte-Back/internal/utils"
	"github.com/jinzhu/copier"
)

func (u *UserService) UserGetByID(id uint) (*schemas.UserResponse, error) {
	user, err := u.UserRepository.UserGetByID(id)
	if err != nil {
		return nil, err
	}

	var userResponse schemas.UserResponse
	_ = copier.Copy(&userResponse, &user)

	return &userResponse, nil
}

func (u *UserService) UserGetByEmail(email string) (*schemas.UserResponse, error) {
	users, err := u.UserRepository.UserGetByEmail(email)
	if err != nil {
		return nil, err
	}	

	var usersResponse schemas.UserResponse
	_ = copier.Copy(&usersResponse, &users)

	return &usersResponse, nil
}

func (u *UserService) UserGetAll() ([]*schemas.UserResponseDTO, error) {
	users, err := u.UserRepository.UserGetAll()
	if err != nil {
		return nil, err
	}

	var usersResponse []*schemas.UserResponseDTO
	_ = copier.Copy(&usersResponse, &users)

	return usersResponse, nil
}

func (u *UserService) UserCreate(userCreate *schemas.UserCreate) (uint, error) {
	return u.UserRepository.UserCreate(userCreate)
}

func (u *UserService) UserUpdate(userUpdate *schemas.UserUpdate) error {
	return u.UserRepository.UserUpdate(userUpdate)
}

func (u *UserService) UserDelete(id uint) error {
	return u.UserRepository.UserDelete(id)
}

func (u *UserService) UserUpdatePassword(userID uint, updatePass *schemas.UserUpdatePassword) error {
	user, err := u.UserRepository.UserGetByID(userID)
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(updatePass.OldPassword, user.Password) {
		return fmt.Errorf("la contraseña actual no es correcta")
	}

	return u.UserRepository.UserUpdatePassword(userID, updatePass)
}