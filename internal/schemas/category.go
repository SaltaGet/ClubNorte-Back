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
	Name string `json:"name" validate:"required" example:"Categoria1"`
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
	return ErrorResponse(422, fmt.Sprintf("error al validar campo(s): %s", errorMessage), err)
}

type CategoryUpdate struct{
	ID   uint   `json:"id" validate:"required" example:"1"`
	Name string `json:"name" validate:"required" example:"Categoria1"`
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
	return ErrorResponse(422, fmt.Sprintf("error al validar campo(s): %s", errorMessage), err)
}