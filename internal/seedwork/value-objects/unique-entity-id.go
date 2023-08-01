package valueobjects

import "github.com/google/uuid"

type UniqueEntityId struct {
	ValueObject[string]
}

func (unique *UniqueEntityId) Validate() error {
	_, error := uuid.Parse(unique.propValue)

	if error != nil {
		return error
	}

	return nil
}

func NewUniqueId(id *string) (*UniqueEntityId, error) {
	var uuidDB string

	if id != nil {
		uuidDB = *id
	} else {
		uuidDB = uuid.New().String()
	}

	valueObject := NewValuableObject[string](uuidDB)

	instance := UniqueEntityId{ValueObject: valueObject}

	error := instance.Validate()

	if error == nil {
		return nil, error
	}

	return &instance, nil
}
