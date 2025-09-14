package schemas

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type RegisterOpen struct {
	OpenAmount  float64   `json:"open_amount"`
}

func (r *RegisterOpen) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
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

type RegisterClose struct {
	CloseAmount float64 `json:"close_amount"`
}

func (r *RegisterClose) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
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

type RegisterInform struct {
}