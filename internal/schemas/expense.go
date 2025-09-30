package schemas

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type ExpenseCreate struct {
	Total         float64 `json:"total" validate:"required"`
	Description   *string `json:"description"`
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

type ExpenseDateRequest struct {
	FromDate string `json:"from_date" example:"2022-01-01"`
	ToDate   string `json:"to_date" example:"2022-12-31"`
}

func (r *ExpenseDateRequest) GetParsedDates() (time.Time, time.Time, error) {
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
