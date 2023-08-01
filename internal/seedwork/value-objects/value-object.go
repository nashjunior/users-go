package valueobjects

import (
	"fmt"
)

type ValueObject[T any] struct {
	propValue T
}

func NewValuableObject[T any](prop T) ValueObject[T] {
	return ValueObject[T]{propValue: prop}
}

func (value *ValueObject[T]) Get() T {
	return value.propValue
}

func (valueObject *ValueObject[T]) ToString() string {
	return fmt.Sprintf("%v", valueObject.propValue)

}
