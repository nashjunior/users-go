package errors

func NotFoundError() *Error {
	return &Error{
		Message: "Item not found",
		Code:    404,
	}
}
