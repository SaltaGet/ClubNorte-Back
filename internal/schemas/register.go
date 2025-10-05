package schemas

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type RegisterOpen struct {
	OpenAmount float64 `json:"open_amount" example:"100.00"`
}

func (r *RegisterOpen) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
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

type RegisterClose struct {
	CloseAmount float64 `json:"close_amount" example:"100.00"`
}

func (r *RegisterClose) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
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

type RegisterInformResponse struct {
	ID uint `json:"id"`
	// PointSale   PointSaleResponse     `json:"point_sale"`
	UserOpen    UserSimpleDTO  `json:"user_open"`
	OpenAmount  float64        `json:"open_amount"`
	HourOpen    time.Time      `json:"hour_open"`
	UserClose   *UserSimpleDTO `json:"user_close"`
	CloseAmount *float64       `json:"close_amount"`
	HourClose   *time.Time     `json:"hour_close"`

	TotalIncomeCash    *float64 `json:"total_income_cash"`
	TotalIncomeOthers  *float64 `json:"total_income_others"`
	TotalExpenseCash   *float64 `json:"total_expense_cash"`
	TotalExpenseOthers *float64 `json:"total_expense_others"`

	IsClose   bool      `json:"is_close"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterInformResponseDTO struct {
	ID uint `json:"id"`

	UserOpen    UserSimpleDTO  `json:"user_open"`
	OpenAmount  float64        `json:"open_amount"`
	HourOpen    time.Time      `json:"hour_open"`
	UserClose   *UserSimpleDTO `json:"user_close"`
	CloseAmount *float64       `json:"close_amount"`
	HourClose   *time.Time     `json:"hour_close"`

	IsClose   bool      `json:"is_close"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterFullResponse struct {
	ID          uint           `json:"id"`
	UserOpen    UserSimpleDTO  `json:"user_open"`
	OpenAmount  float64        `json:"open_amount"`
	HourOpen    time.Time      `json:"hour_open"`
	UserClose   *UserSimpleDTO `json:"user_close"`
	CloseAmount *float64       `json:"close_amount"`
	HourClose   *time.Time     `json:"hour_close"`

	TotalIncomeCash    *float64 `json:"total_income_cash"`
	TotalIncomeOthers  *float64 `json:"total_income_others"`
	TotalExpenseCash   *float64 `json:"total_expense_cash"`
	TotalExpenseOthers *float64 `json:"total_expense_others"`

	IsClose   bool      `json:"is_close"`
	CreatedAt time.Time `json:"created_at"`

	Income             []IncomeSimpleResponse          `json:"income"`
	IncomeSportsCourts []IncomeSportsCourtsResponseDTO `json:"income_sports_courts"`
	Expense           []ExpenseSimpleResponseDTO      `json:"expenses"`
}
