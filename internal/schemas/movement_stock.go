package schemas

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type MovementStockResponse struct {
	ID          uint            `json:"id"`
	User        UserResponse    `json:"user"`
	Product     ProductResponse `json:"product"`
	Amount      float64         `json:"amount"`
	FromID      uint            `json:"from_id"`
	FromType    string          `json:"from_type"`
	ToID        uint            `json:"to_id"`
	ToType      string          `json:"to_type"`
	IgnoreStock bool            `json:"ignore_stock"`
	CreatedAt   time.Time       `json:"created_at"`
}

type MovementStockResponseDTO struct {
	ID          uint                     `json:"id"`
	User        UserResponse             `json:"user"`
	Product     ProductSimpleResponseDTO `json:"product"`
	Amount      float64                  `json:"amount"`
	FromID      uint                     `json:"from_id"`
	FromType    string                   `json:"from_type"`
	ToID        uint                     `json:"to_id"`
	ToType      string                   `json:"to_type"`
	IgnoreStock bool                     `json:"ignore_stock"`
	CreatedAt   time.Time                `json:"created_at"`
}

type MovementStock struct {
	ProductID uint    `json:"product_id" validate:"required"`
	Amount    float64 `json:"amount" validate:"required"`

	FromType string `json:"from_type" validate:"oneof=deposit point_sale"`
	FromID   uint   `json:"from_id" validate:"required"`

	ToType string `json:"to_type" validate:"oneof=deposit point_sale"`
	ToID   uint   `json:"to_id" validate:"required"`

	IgnoreStock bool `json:"ignore_stock" validate:"required"`
}

func (m *MovementStock) Validate() error {
	validate := validator.New()
	err := validate.Struct(m)
	if err == nil {
		return nil
	}

	validatorErr := err.(validator.ValidationErrors)[0]
	field := validatorErr.Field()
	tag := validatorErr.Tag()
	params := validatorErr.Param()

	errorMessage := field + " " + tag + " " + params
	return fmt.Errorf("error al validar el login: %s", errorMessage)
}
