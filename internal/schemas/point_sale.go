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
	Name        string `json:"name" validate:"required"`
	Description *string `json:"description"`
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
	return fmt.Errorf("error al validar el login: %s", errorMessage)
}

type PointSaleUpdate struct {
	ID          uint   `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description *string `json:"description"`
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
	return fmt.Errorf("error al validar el login: %s", errorMessage)
}