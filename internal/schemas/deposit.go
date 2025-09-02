package schemas

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type DepositResponse struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description *string `json:"description"`
	Category    CategoryResponse `json:"category"`
	Price       float64 `json:"price"`
	Stock       float64    `json:"stock"`
}

type DepositUpdateStock struct {
	ProductID uint   `json:"product_id"`
	Stock     float64 `json:"stock"`
	Method    string `json:"method" validate:"oneof=add subtract set" example:"add|subtract|set"`
}

func (d *DepositUpdateStock) Validate() error {
	validate := validator.New()
	err := validate.Struct(d)
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