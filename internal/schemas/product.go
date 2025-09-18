package schemas

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ProductFullResponse struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    CategoryResponse `json:"category"`
	Price       float64 `json:"price"`
	StockPointSales []*PointSaleStock `gorm:"foreignKey:ProductID" json:"stock_point_sales"`
	StockDeposit   *StockDepositResponse   `gorm:"foreignKey:ProductID" json:"stock_deposit"`
	Notifier    bool     `json:"notifier"`
	MinAmount   float64  `json:"min_amount"`
}

type ProductResponse struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    CategoryResponse `json:"category"`
	Price       float64 `json:"price"`
	Stock       float64    `json:"stock"`
	Notifier    bool     `json:"notifier"`
	MinAmount   float64  `json:"min_amount"`
}

type ProductResponseDTO struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Category    *CategoryResponse `json:"category,omitempty"`
	Price       float64 `json:"price"`
	Stock       float64    `json:"stock"`
	Notifier    bool     `json:"notifier"`
	MinAmount   float64  `json:"min_amount"`
}

type ProductSimpleResponse struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Price       float64 `json:"price"`
	Stock       float64    `json:"stock"`
	Notifier    bool     `json:"notifier"`
	MinAmount   float64  `json:"min_amount"`
}

type ProductSimpleResponseDTO struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Price       float64 `json:"price"`
}

type ProductCreate struct {
	Code        string  `json:"code" validate:"required" example:"ABC123"`
	Name        string  `json:"name" validate:"required" example:"Producto1"`
	Description *string `json:"description"`
	CategoryID  uint    `json:"category_id" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Notifier    bool     `json:"notifier"`
	MinAmount   float64  `json:"min_amount"`
}

func (p *ProductCreate) Validate () error {
	validate := validator.New()
	err := validate.Struct(p)
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

type ProductUpdate struct {
	ID          uint    `json:"id" validate:"required"`
	Code        string  `json:"code" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	CategoryID  uint    `json:"category_id" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Notifier    bool     `json:"notifier"`
	MinAmount   float64  `json:"min_amount"`
}

func (p *ProductUpdate) Validate () error {
	validate := validator.New()
	err := validate.Struct(p)
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