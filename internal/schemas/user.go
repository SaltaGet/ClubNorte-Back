package schemas

import (
	"fmt"
	"regexp"
	"strings"

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
	FirstName     string  `json:"first_name" validate:"required" example:"John"`
	LastName      string  `json:"last_name" validate:"required" example:"Doe"`
	Address       *string `json:"address,omitempty" example:"address|null"`
	Cellphone     *string `json:"cellphone,omitempty" example:"cellphone|null"`
	Email         string  `json:"email" validate:"email,required" example:"a@b.com"`
	Username      string  `json:"username" validate:"required" example:"johndoe"`
	Password      string  `json:"password" validate:"required,password" example:"Password123*"`
	RoleID        uint    `json:"role_id" validate:"required" example:"1"`
	PointSaleIDs []uint  `json:"point_sales_ids" validate:"required" example:"1,2,3"`
}

func validatePassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()

    if len(password) < 8 {
        return false
    }
    if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
        return false
    }
    if !regexp.MustCompile(`[0-9]`).MatchString(password) {
        return false
    }
    if !regexp.MustCompile(`[^A-Za-z0-9]`).MatchString(password) {
        return false
    }
    return true
}


// func (u *UserCreate) Validate() error {
// 	validate := validator.New()
// 	validate.RegisterValidation("password", validatePassword)
// 	err := validate.Struct(u)
// 	if err == nil {
// 		return nil
// 	}

// 	validatorErr := err.(validator.ValidationErrors)[0]
// 	field := validatorErr.Field()
// 	tag := validatorErr.Tag()
// 	params := validatorErr.Param()

// 	errorMessage := field + " " + tag + " " + params
// 	return ErrorResponse(422, fmt.Sprintf("error al validar campo(s): %s", errorMessage), err)
// }
func (u *UserCreate) Validate() error {
    validate := validator.New()
    validate.RegisterValidation("password", validatePassword)

    err := validate.Struct(u)
    if err == nil {
        return nil
    }

    validatorErr := err.(validator.ValidationErrors)[0]
    field := validatorErr.Field()
    tag := validatorErr.Tag()

    var errorMessage string
    switch tag {
    case "required":
        errorMessage = fmt.Sprintf("el campo %s es obligatorio", field)
    case "email":
        errorMessage = fmt.Sprintf("el campo %s debe ser un email válido", field)
    case "password":
        errorMessage = "el campo password debe tener al menos 8 caracteres, una letra mayúscula, un número y un carácter especial"
    default:
        errorMessage = fmt.Sprintf("el campo %s no cumple la validación %s", field, tag)
    }

    return ErrorResponse(422, fmt.Sprintf("error al validar campo(s): %s", errorMessage), err)
}


type UserUpdate struct {
	ID            uint    `json:"id" validate:"required" example:"1"`
	FirstName     string  `json:"first_name" validate:"required" example:"John"`
	LastName      string  `json:"last_name" validate:"required" example:"Doe"`
	Address       *string `json:"address,omitempty" example:"address|null"`
	Cellphone     *string `json:"cellphone,omitempty" example:"cellphone|null"`
	Email         string  `json:"email" validate:"email,required" example:"a@b.com"`
	Username      string  `json:"username" validate:"required" example:"johndoe"`
	RoleID        uint    `json:"role_id" validate:"required" example:"1"`
	IsActive      bool    `json:"is_active" validate:"required" example:"true"`
	PointSaleIDs []uint  `json:"point_sales_ids" validate:"required" example:"1,2,3"`
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
	OldPassword   string  `json:"old_password" validate:"required,password" example:"Password123*"`
	NewPassword   string  `json:"new_password" validate:"required,password" example:"Password123*"`
	ConfirmPass   string  `json:"confirm_pass" validate:"required,password" example:"Password123*"`
}

func (u *UserUpdatePassword) Validate() error {
    validate := validator.New()
    validate.RegisterValidation("password", validatePassword) // registrar antes de Struct()

    // validación de campos con reglas
    err := validate.Struct(u)
    if err != nil {
        var messages []string
        for _, e := range err.(validator.ValidationErrors) {
            field := e.Field()
            tag := e.Tag()

            var msg string
            switch tag {
            case "required":
                msg = fmt.Sprintf("el campo %s es obligatorio", field)
            case "password":
                msg = fmt.Sprintf("el campo %s debe tener al menos 8 caracteres, una mayúscula, un número y un caracter especial", field)
            default:
                msg = fmt.Sprintf("el campo %s no cumple la validación %s", field, tag)
            }
            messages = append(messages, msg)
        }
        return ErrorResponse(422, strings.Join(messages, ", "), err)
    }

    // validación manual: confirmación de contraseña
    if u.NewPassword != u.ConfirmPass {
        return ErrorResponse(422, "las contraseñas no coinciden", nil)
    }

    return nil
}

// func (u *UserUpdatePassword) Validate() error {
// 	validate := validator.New()
// 	err := validate.Struct(u)

// 	validate.RegisterValidation("password", validatePassword)

// 	if u.NewPassword != u.ConfirmPass {
// 		return fmt.Errorf("las contraseñas no coinciden")
// 	}

// 	if err == nil {
// 		return nil
// 	}

// 	validatorErr := err.(validator.ValidationErrors)[0]
// 	field := validatorErr.Field()
// 	tag := validatorErr.Tag()
// 	params := validatorErr.Param()

// 	errorMessage := field + " " + tag + " " + params
// 	return ErrorResponse(422, fmt.Sprintf("error al validar campo(s): %s", errorMessage), err)
// }