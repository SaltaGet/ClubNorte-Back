package schemas

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type Totals struct {
	Cash   float64 `json:"cash"`
	Others float64 `json:"others"`
}

type IncomeCreate struct {
	Items         []IncomeItemCreate `json:"items" validate:"required,dive"`
	Description   *string            `json:"description" example:"description|null"`
	PaymentMethod string             `json:"payment_method" validate:"required,oneof=efectivo tarjeta transferencia" example:"efectivo|tarjeta|transferencia"`
}

func (i *IncomeCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(i)
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

type IncomeUpdate struct {
	ID            uint               `json:"id" validate:"required" example:"1"`
	Items         []IncomeItemCreate `json:"items" validate:"required,dive"`
	Description   *string            `json:"description" example:"description|null"`
	PaymentMethod string             `json:"payment_method" validate:"required,oneof=efectivo tarjeta transferencia" example:"efectivo|tarjeta|transferencia"`
}

func (i *IncomeUpdate) Validate() error {
	validate := validator.New()
	err := validate.Struct(i)
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

type IncomeItemCreate struct {
	ProductID uint    `json:"product_id" validate:"required" example:"1"`
	Quantity  float64     `json:"quantity" validate:"required,gt=0" example:"10"`
}

type IncomeResponse struct {
	ID            uint                 `json:"id"`
	UserResponse  UserSimpleDTO        `json:"user"`
	Items         []IncomeItemResponse `json:"items"`
	Description   *string              `json:"description"`
	Total         float64              `json:"total"`
	PaymentMethod string               `json:"payment_method"`
	CreatedAt     time.Time            `json:"created_at"`
}

type IncomeSimpleResponse struct {
	ID            uint                 `json:"id"`
	Items         []IncomeItemResponse `json:"items"`
	Description   *string              `json:"description"`
	Total         float64              `json:"total"`
	PaymentMethod string               `json:"payment_method"`
	CreatedAt     time.Time            `json:"created_at"`
}

type IncomeItemResponse struct {
	ID       uint                     `json:"id"`
	Product  ProductSimpleResponseDTO `json:"product"`
	Quantity float64                  `json:"quantity"`
	Price    float64                  `json:"price"`
	Subtotal float64                  `json:"subtotal"`
}

type IncomeResponseDTO struct {
	ID            uint          `json:"id"`
	UserResponse  UserSimpleDTO `json:"user"`
	Total         float64       `json:"total"`
	PaymentMethod string        `json:"payment_method"`
	CreatedAt     time.Time     `json:"created_at"`
}
