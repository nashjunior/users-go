package validators

func MakeValidator() AbstractValidator[any] {
	validation := MakeV10Validator[any]()

	return validation.AbstractValidator
}
