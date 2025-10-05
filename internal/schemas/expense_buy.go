package schemas

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type ExpenseBuyResponse struct {
	ID              uint                      `json:"id"`
	User            UserSimpleDTO             `json:"user"`
	Description     *string                   `json:"description"`
	ItemExpenseBuys []*ItemExpenseBuyResponse `json:"item_expense_buys"`
	PaymentMethod   string                    `json:"payment_method"`
	CreatedAt       time.Time                 `json:"created_at"`
	Total           float64                   `json:"total"`
}

type ExpenseBuyResponseDTO struct {
	ID            uint          `json:"id"`
	User          UserSimpleDTO `json:"user"`
	PaymentMethod string        `json:"payment_method"`
	Total         float64       `json:"total"`
	CreatedAt     time.Time     `json:"created_at"`
}

type ItemExpenseBuyResponse struct {
	ID        uint                     `json:"id"`
	Product   ProductSimpleResponseDTO `json:"product"`
	Quantity  uint                     `json:"quantity"`
	Price     float64                  `json:"price"`
	Subtotal  float64                  `json:"subtotal"`
	CreatedAt time.Time                `json:"created_at"`
}

type ExpenseBuyCreate struct {
	Description     *string           `json:"description" example:"description|null"`
	ItemExpenseBuys []*ItemExpenseBuyCreate `json:"item_expense_buys" validate:"required,dive"`
	PaymentMethod   string            `json:"payment_method" validate:"oneof=efectivo tarjeta transferencia" example:"efectivo|tarjeta|transferencia"`
}

type ItemExpenseBuyCreate struct {
	ProductID uint `json:"product_id" validate:"required" example:"1"`
	Quantity  float64 `json:"quantity" validate:"required" example:"10"`
	Price     float64 `json:"price" validate:"required" example:"100"`
}

func (i *ExpenseBuyCreate) Validate() error {
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
