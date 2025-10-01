package schemas

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type PointSaleContext struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PointSaleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Description *string `json:"description"`
}

type PointSaleCreate struct {
	Name        string `json:"name" validate:"required" example:"Punto de venta 1"`
	Description *string `json:"description" example:"description|null"`
}

func (p *PointSaleCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)
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

type PointSaleUpdate struct {
	ID          uint   `json:"id" validate:"required" example:"1"`
	Name        string `json:"name" validate:"required" example:"Punto de venta 1"`
	Description *string `json:"description" example:"description|null"`
}

func (p *PointSaleUpdate) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)
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