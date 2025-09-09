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
	role, err := u.RoleRepository.RoleGetByID(userCreate.RoleID)
	if err != nil {
		return 0, err
	} 

	if role.Name == "admin" {
		return 0, schemas.ErrorResponse(400, "el usuario no puede ser creado con este rol", nil)
	}

	return u.UserRepository.UserCreate(userCreate)
}

func (u *UserService) UserUpdate(userUpdate *schemas.UserUpdate) error {
	role, err := u.RoleRepository.RoleGetByID(userUpdate.RoleID)
	if err != nil {
		return err
	} 

	if role.Name == "admin" {
		return schemas.ErrorResponse(400, "el usuario no puede ser editado con este rol", fmt.Errorf("el usuario no puede ser editado con este rol %s", role.Name))
	}

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
		return schemas.ErrorResponse(400, "Contraseña incorrecta", fmt.Errorf("la contraseña actual no es correcta"))
	}

	return u.UserRepository.UserUpdatePassword(userID, updatePass)
}