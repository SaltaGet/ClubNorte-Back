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
	Description     *string           `json:"description"`
	ItemExpenseBuys []*ItemExpenseBuyCreate `json:"item_expense_buys" validate:"required,dive"`
	PaymentMethod   string            `json:"payment_method" validate:"oneof=efectivo tarjeta transferencia"`
}

type ItemExpenseBuyCreate struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  float64 `json:"quantity" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
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

type ExpenseBuyDateRequest struct {
	FromDate string `json:"from_date" example:"2022-01-01"`
	ToDate   string `json:"to_date" example:"2022-12-31"`
}

func (r *ExpenseBuyDateRequest) GetParsedDates() (time.Time, time.Time, error) {
	loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")

	from, err := time.ParseInLocation("2006-01-02", r.FromDate, loc)
	if err != nil {
		return time.Time{}, time.Time{}, ErrorResponse(422, "error al parsear la fecha de inicio", err)
	}

	to, err := time.ParseInLocation("2006-01-02", r.ToDate, loc)
	if err != nil {
		return time.Time{}, time.Time{}, ErrorResponse(422, "error al parsear la fecha de fin", err)
	}

	from = time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, from.Location())
	to = time.Date(to.Year(), to.Month(), to.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), to.Location())

	return from, to, nil
}