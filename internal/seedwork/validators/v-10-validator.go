package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type V10Validator[T any] struct {
	AbstractValidator[T]
}

func getErrorMsgByTag(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("The field %s is required", e.Field())

	case "max":
		return fmt.Sprintf("The field %s is at most %s", e.Field(), e.Param())

	case "min":
		return fmt.Sprintf("The field %s is at least %s", e.Field(), e.Param())

	default:
		return "A default error message"
	}
}

func (v10Validator *V10Validator[T]) Validate(data any) bool {
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		v10Validator.AbstractValidator.Errors = &map[string][]string{}

		for _, errFields := range err.(validator.ValidationErrors) {
			errorsMap := *v10Validator.AbstractValidator.Errors
			errorMessage := getErrorMsgByTag(errFields)
			errorsMap[errFields.Field()] = append(errorsMap[errFields.Field()], errorMessage)
			*v10Validator.AbstractValidator.Errors = errorsMap
		}

		return false
	} else {
		v10Validator.ValidatedData = data.(T)
	}

	return true
}

func MakeV10Validator[T any]() V10Validator[T] {
	return V10Validator[T]{}
}
