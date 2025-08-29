package schemas

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Login struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

func(l *Login) Validate() error {
	validate := validator.New()
	err := validate.Struct(l)
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
