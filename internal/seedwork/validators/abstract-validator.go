package validators

type IErrorField = map[string][]string

type IValidator interface {
	Validate(data any) bool
}

type AbstractValidator[T any] struct {
	Errors        *IErrorField
	ValidatedData T
	IValidator
}
