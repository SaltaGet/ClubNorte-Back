package schemas

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type IncomeSportsCourtsCreate struct {
	SportsCourtID        uint      `json:"sports_court_id" validate:"required"`
	Shift                string    `json:"shift" validate:"required,oneof=mañana tarde noche"`
	DatePlay             time.Time `json:"date_play" validate:"required"`
	PartialPay           float64   `json:"partial_pay"`
	PartialPaymentMethod string    `json:"partial_payment_method" validate:"oneof=efectivo tarjeta transferencia"`
	DatePartialPay       time.Time `json:"date_partial_pay"`
	Price                float64   `json:"price"`

	// RestPay           float64   `json:"rest_pay"`
	// RestPaymentMethod string    `json:"rest_payment_method" validate:"oneof=efectivo tarjeta transferencia"`
	// DateRestPay       time.Time `json:"date_rest_pay"`
	// RestRegisterID    *uint     `json:"rest_register_id"`
	//GAS string `Habia pensado en un sistema experto para recomendar estrategias didácticas. La idea es que el profesor debe ingresar un tema(ej: trigonometría,sucesiones, movimiento parabólico, etc) y el sistema sugiere un tipo de recurso (juego, proyecto, software). REGLAS: si el tema es funciones -> uso de GeoGebra. Si el tema es probabilidades y estadística -> juegos de dados o cartas. También había pensado en algo para identificar dificultades de los estudiantes y recomendaciones de ejercicios o quizás videos de internet. Por las dudas estoy por recibirme de prof de Secundaria en matemáticas. Tengo muchas ideas y no me decido `
}

func (i *IncomeSportsCourtsCreate) Validate() error {
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

type IncomeSportsCourtsUpdate struct {
	ID                   uint       `json:"id" validate:"required"`
	SportsCourtID        uint       `json:"sports_court_id" validate:"required"`
	Shift                string     `json:"shift" validate:"required,oneof=mañana tarde noche"`
	DatePlay             time.Time  `json:"date_play" validate:"required"`
	PartialPay           *float64   `json:"partial_pay"`
	PartialPaymentMethod *string    `json:"partial_payment_method" validate:"oneof=efectivo tarjeta transferencia"`
	DatePartialPay       *time.Time `json:"date_partial_pay"`
	PartialRegisterID    *uint      `json:"partial_register_id"`

	RestPay           *float64   `json:"rest_pay"`
	RestPaymentMethod *string    `json:"rest_payment_method" validate:"oneof=efectivo tarjeta transferencia"`
	DateRestPay       *time.Time `json:"date_rest_pay"`
	RestRegisterID    *uint      `json:"rest_register_id"`

	Price        *float64 `json:"price"`
}

func (i *IncomeSportsCourtsUpdate) Validate() error {
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

type IncomeSportsCourtsRestPay struct {
	ID                uint    `json:"id" validate:"required"`
	RestPay           float64 `json:"rest_pay"`
	RestPaymentMethod string  `json:"rest_payment_method" validate:"oneof=efectivo tarjeta transferencia"`
}

func (i *IncomeSportsCourtsRestPay) Validate() error {
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

type IncomeSportsCourtsResponse struct {
	ID            uint                 `json:"id"`
	UserResponse  UserSimpleDTO        `json:"user"`
	Items         []IncomeItemResponse `json:"items"`
	Description   *string              `json:"description"`
	Total         float64              `json:"total"`
	PaymentMethod string               `json:"payment_method"`
	CreatedAt     time.Time            `json:"created_at"`
}

type IncomeSportsCourtsResponseDTO struct {
	ID            uint          `json:"id"`
	UserResponse  UserSimpleDTO `json:"user"`
	Total         float64       `json:"total"`
	PaymentMethod string        `json:"payment_method"`
	CreatedAt     time.Time     `json:"created_at"`
}

type IncomeSportsCourtsDateRequest struct {
	FromDate string `json:"from_date" example:"2022-01-01"`
	ToDate   string `json:"to_date" example:"2022-12-31"`
}

func (r *IncomeSportsCourtsDateRequest) GetParsedDates() (time.Time, time.Time, error) {
	loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")

	from, err := time.ParseInLocation("2006-01-02", r.FromDate, loc)
	if err != nil {
		return time.Time{}, time.Time{}, ErrorResponse(422, "error al parsear la fecha de inicio", err)
	}

	to, err := time.ParseInLocation("2006-01-02", r.ToDate, loc)
	if err != nil {
		return time.Time{}, time.Time{}, ErrorResponse(422, "error al parsear la fecha de fin", err)
	}

	// Ajustar horas
	from = time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, from.Location())
	to = time.Date(to.Year(), to.Month(), to.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), to.Location())

	return from, to, nil
}
