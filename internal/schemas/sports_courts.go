package schemas

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type SportCourtResponse struct {
	ID          uint                `json:"id"`
	Code        string              `json:"code"`
	Name        string              `json:"name"`
	Description *string             `json:"description,omitempty"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	PointSales  []PointSaleResponse `json:"point_sales"`
}

type SportCourtResponseDTO struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SportCourtCreate struct {
	Code         string  `json:"code" validate:"required" example:"ABC123"`
	Name         string  `json:"name" validate:"required" example:"Pista 1"`
	Description  *string `json:"description,omitempty" example:"description|null"`
}

func (s *SportCourtCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
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

type SportCourtUpdate struct {
	ID           uint    `json:"id" validate:"required" example:"1"`
	Code         string  `json:"code" validate:"required" example:"ABC123"`
	Name         string  `json:"name" validate:"required" example:"Pista 1"`
	Description  *string `json:"description,omitempty" example:"description|null"`
}

func (s *SportCourtUpdate) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
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
