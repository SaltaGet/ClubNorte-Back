package schemas

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type UserResponse struct {
	ID        uint         `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Address   *string      `json:"address"`
	Cellphone *string      `json:"cellphone"`
	Email     string       `json:"email"`
	Username  string       `json:"username"`
	IsActive  bool         `json:"is_active"`
	IsAdmin   bool         `json:"is_admin"`
	Role      RoleResponse `json:"role"`
	PointSales []PointSaleResponse  `json:"point_sales"`
}

type UserResponseDTO struct {
	ID        uint         `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Address   *string      `json:"address"`
	Cellphone *string      `json:"cellphone"`
	Email     string       `json:"email"`
	Username  string       `json:"username"`
	IsActive  bool         `json:"is_active"`
	IsAdmin   bool         `json:"is_admin"`
	Role      RoleResponse `json:"role"`
}

type UserSimpleDTO struct {
	ID        uint         `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Address   *string      `json:"address"`
	Cellphone *string      `json:"cellphone"`
	Email     string       `json:"email"`
	Username  string       `json:"username"`
}

type UserResponseToken struct {
	ID        uint         `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Address   *string      `json:"address"`
	Cellphone *string      `json:"cellphone"`
	Email     string       `json:"email"`
	Username  string       `json:"username"`
	IsAdmin   bool         `json:"is_admin"`
	Role      RoleResponse `json:"role"`
}

type UserContext struct {
	ID        uint    `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Address   *string `json:"address"`
	Cellphone *string `json:"cellphone"`
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	IsAdmin   bool    `json:"is_admin"`
	IsActive  bool    `json:"is_active"`
	RoleID    uint    `json:"role_id"`
	Role      string  `json:"role"`
}

type UserCreate struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Address       *string `json:"address,omitempty"`
	Cellphone     *string `json:"cellphone,omitempty"`
	Email         string  `json:"email"`
	Username      string  `json:"username"`
	Password      string  `json:"password"`
	RoleID        uint    `json:"role_id"`
	PointSaleIDs []uint  `json:"point_sales_ids"`
}


func (u *UserCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err == nil {
		return nil
	}

	validatorErr := err.(validator.ValidationErrors)[0]
	field := validatorErr.Field()
	tag := validatorErr.Tag()
	params := validatorErr.Param()

	errorMessage := field + " " + tag + " " + params
	return ErrorResponse(422, fmt.Sprintf("error al validar campo(s): %s", errorMessage), err)
}

type UserUpdate struct {
	ID            uint    `json:"id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Address       *string `json:"address,omitempty"`
	Cellphone     *string `json:"cellphone,omitempty"`
	Email         string  `json:"email"`
	Username      string  `json:"username"`
	RoleID        uint    `json:"role_id"`
	IsActive      bool    `json:"is_active"`
	PointSaleIDs []uint  `json:"point_sales_ids"`
}


func (u *UserUpdate) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err == nil {
		return nil
	}

	validatorErr := err.(validator.ValidationErrors)[0]
	field := validatorErr.Field()
	tag := validatorErr.Tag()
	params := validatorErr.Param()

	errorMessage := field + " " + tag + " " + params
	return ErrorResponse(422, fmt.Sprintf("error al validar campo(s): %s", errorMessage), err)
}

type UserUpdatePassword struct {
	OldPassword   string  `json:"old_password"`
	NewPassword   string  `json:"new_password"`
	ConfirmPass   string  `json:"confirm_pass"`
}


func (u *UserUpdatePassword) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)

	if u.NewPassword != u.ConfirmPass {
		return fmt.Errorf("las contraseñas no coinciden")
	}

	if err == nil {
		return nil
	}

	validatorErr := err.(validator.ValidationErrors)[0]
	field := validatorErr.Field()
	tag := validatorErr.Tag()
	params := validatorErr.Param()

	errorMessage := field + " " + tag + " " + params
	return ErrorResponse(422, fmt.Sprintf("error al validar campo(s): %s", errorMessage), err)
}