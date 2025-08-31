package schemas

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CategoryResponse struct{
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CategoryCreate struct{
	Name string `json:"name" validate:"required"`
}	

func (c CategoryCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err == nil {
		return nil
	}

	validatorErr := err.(validator.ValidationErrors)[0]
	field := validatorErr.Field()
	tag := validatorErr.Tag()
	params := validatorErr.Param()

	errorMessage := field + " " + tag + " " + params
	return fmt.Errorf("error al validar el login: %s", errorMessage)
}

type CategoryUpdate struct{
	ID   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

func (c CategoryUpdate) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err == nil {
		return nil
	}

	validatorErr := err.(validator.ValidationErrors)[0]
	field := validatorErr.Field()
	tag := validatorErr.Tag()
	params := validatorErr.Param()

	errorMessage := field + " " + tag + " " + params
	return fmt.Errorf("error al validar el login: %s", errorMessage)
}