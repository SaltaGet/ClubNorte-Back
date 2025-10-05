package schemas

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type IncomeSportsCourtsCreate struct {
	SportsCourtID        uint      `json:"sports_court_id" validate:"required" example:"1"`
	Shift                string    `json:"shift" validate:"required,oneof=mañana tarde noche" example:"mañana|tarde|noche"`
	DatePlay             time.Time `json:"date_play" validate:"required" example:"2023-08-01T12:00:00Z"`
	PartialPay           float64   `json:"partial_pay" example:"100.00"`
	PartialPaymentMethod string    `json:"partial_payment_method" validate:"oneof=efectivo tarjeta transferencia" example:"efectivo|tarjeta|transferencia"`
	Total                float64   `json:"price" example:"200.00"`

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

// type IncomeSportsCourtsUpdate struct {
// 	ID                   uint       `json:"id" validate:"required" example:"1"`
// 	SportsCourtID        uint       `json:"sports_court_id" validate:"required" example:"1"`
// 	Shift                string     `json:"shift" validate:"required,oneof=mañana tarde noche" example:"mañana|tarde|noche"`
// 	DatePlay             time.Time  `json:"date_play" validate:"required" example:"2023-08-01T12:00:00Z"`
// 	PartialPay           *float64   `json:"partial_pay" example:"100.00"`
// 	PartialPaymentMethod *string    `json:"partial_payment_method" validate:"oneof=efectivo tarjeta transferencia" example:"efectivo|tarjeta|transferencia"`
// 	DatePartialPay       *time.Time `json:"date_partial_pay" example:"2023-08-01T12:00:00Z"`
// 	PartialRegisterID    *uint      `json:"partial_register_id" example:"1"`

// 	RestPay           *float64   `json:"rest_pay" example:"100.00"`
// 	RestPaymentMethod *string    `json:"rest_payment_method" validate:"oneof=efectivo tarjeta transferencia" example:"efectivo|tarjeta|transferencia"`
// 	DateRestPay       *time.Time `json:"date_rest_pay" example:"2023-08-01T12:00:00Z"`
// 	RestRegisterID    *uint      `json:"rest_register_id" example:"1"`

// 	Price        *float64 `json:"price"`
// }

// func (i *IncomeSportsCourtsUpdate) Validate() error {
// 	validate := validator.New()
// 	err := validate.Struct(i)
// 	if err == nil {
// 		return nil
// 	}

// 	validatorErr := err.(validator.ValidationErrors)[0]
// 	field := validatorErr.Field()
// 	tag := validatorErr.Tag()
// 	params := validatorErr.Param()

// 	errorMessage := field + " " + tag + " " + params
// 	return ErrorResponse(422, fmt.Sprintf("error al validar campo(s): %s", errorMessage), err)
// }

type IncomeSportsCourtsRestPay struct {
	ID                uint    `json:"id" validate:"required" example:"1"`
	RestPay           float64 `json:"rest_pay" example:"100.00"`
	RestPaymentMethod string  `json:"rest_payment_method" validate:"oneof=efectivo tarjeta transferencia" example:"efectivo|tarjeta|transferencia"` 
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
	ID           uint          `json:"id"`
	UserResponse UserSimpleDTO `json:"user"`
	Description  *string       `json:"description"`

	PartialPay           float64   `json:"partial_pay"`
	PartialPaymentMethod string    `json:"partial_payment_method"`
	DatePartialPay       time.Time `json:"date_partial_pay"`

	RestPay           *float64   `json:"rest_pay"`
	RestPaymentMethod *string    `json:"rest_payment_method"`
	DateRestPay       *time.Time `json:"date_rest_pay"`

	Total     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type IncomeSportsCourtsResponseDTO struct {
	ID           uint          `json:"id"`
	Description  *string       `json:"description"`

	PartialPay           float64   `json:"partial_pay"`
	PartialPaymentMethod string    `json:"partial_payment_method"`
	DatePartialPay       time.Time `json:"date_partial_pay"`

	RestPay           *float64   `json:"rest_pay"`
	RestPaymentMethod *string    `json:"rest_payment_method"`
	DateRestPay       *time.Time `json:"date_rest_pay"`

	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

