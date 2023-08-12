package errors

import "fmt"

type Error struct {
	Message interface{}
	Code    int
}

// Error implements the error interface for the CustomError type.
func (e *Error) Error() string {
	return fmt.Sprintf("Error Code: %d - %s", e.Code, e.Message)
}
