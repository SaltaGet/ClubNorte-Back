package schemas

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type ExpenseCreate struct {
	Total         float64 `json:"total" validate:"required" example:"1000"`
	Description   *string `json:"description" example:"description|null"`
	PaymentMethod string  `json:"payment_method" validate:"required,oneof=efectivo tarjeta transferencia"`
}

func (i *ExpenseCreate) Validate() error {
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

type ExpenseResponse struct {
	ID            uint      `json:"id"`
	// PointSale     PointSaleResponse `json:"point_sale"`
	User          UserSimpleDTO      `json:"user"`
	RegisterID    uint      `json:"register_id"`
	Description   *string   `json:"description"`
	Total         float64   `json:"total"`
	PaymentMethod string    `json:"payment_method"`
	CreatedAt     time.Time `json:"created_at"`
}

type ExpenseResponseDTO struct {
	ID            uint          `json:"id"`
	User  UserSimpleDTO `json:"user"`
	Total         float64       `json:"total"`
	PaymentMethod string        `json:"payment_method"`
	CreatedAt     time.Time     `json:"created_at"`
}

type ExpenseSimpleResponseDTO struct {
	ID            uint          `json:"id"`
	Total         float64       `json:"total"`
	PaymentMethod string        `json:"payment_method"`
	CreatedAt     time.Time     `json:"created_at"`
}

