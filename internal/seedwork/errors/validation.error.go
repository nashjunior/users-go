package errors

import "school-api/internal/seedwork/validators"

func ValidationError(errors validators.IErrorField) *Error {
	return &Error{
		Message: "Entity Validation Error",
		Code:    404,
	}
}
