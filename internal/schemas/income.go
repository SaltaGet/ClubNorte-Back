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
	Items         []IncomeItemCreate `json:"items" validate:"required,dive"` // <- dive recorre cada elemento
	Description   *string            `json:"description"`
	PaymentMethod string             `json:"payment_method" validate:"required,oneof=efectivo tarjeta transferencia"`
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
	ID            uint               `json:"id" validate:"required"`
	Items         []IncomeItemCreate `json:"items" validate:"required,dive"` // <- dive recorre cada elemento
	Description   *string            `json:"description"`
	PaymentMethod string             `json:"payment_method" validate:"required,oneof=efectivo tarjeta transferencia"`
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
	ProductID uint    `json:"product_id" validate:"required"`
	Quantity  float64     `json:"quantity" validate:"required,gt=0"`
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

type IncomeDateRequest struct {
	FromDate string `json:"from_date" example:"2022-01-01"`
	ToDate   string `json:"to_date" example:"2022-12-31"`
}

func (r *IncomeDateRequest) GetParsedDates() (time.Time, time.Time, error) {
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