package errors

func InvalidUUIDError() *Error {
	messsage := "Must be a valid UUID"
	return &Error{
		Message: messsage,
		Code:    404,
	}
}
