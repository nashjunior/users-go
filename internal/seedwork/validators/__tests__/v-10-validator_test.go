package validators

import (
	"school-api/internal/seedwork/validators"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StubV10Validator struct {
	validators.V10Validator[any]
	mock.Mock
}

func (stubValidator *StubV10Validator) Validate(data any) bool {
	stubValidator.V10Validator.AbstractValidator.Errors = &map[string][]string{
		"field": []string{"required"},
	}

	args := stubValidator.Called(data)

	return args.Bool(0)
}

type StubV10ValidatorValidated struct {
	validators.V10Validator[any]
	mock.Mock
}

func (stubValidator *StubV10ValidatorValidated) Validate(data any) bool {
	stubValidator.V10Validator.AbstractValidator.ValidatedData = data

	args := stubValidator.Called(data)

	return args.Bool(0)
}

func TestV10ValidatorInitializationErrorsAndValidatedDataNil(t *testing.T) {

	stub := StubV10Validator{
		V10Validator: validators.V10Validator[any]{},
	}

	assert.Nil(t, stub.AbstractValidator.Errors)
	assert.Nil(t, stub.AbstractValidator.ValidatedData)

}

func TestV10ValidatorWithErrors(t *testing.T) {
	testObj := new(StubV10Validator)
	structure := struct{ Field string }{Field: "some value"}

	testObj.On("Validate", structure).Return(false)
	testObj.Validate(structure)

	testObj.AssertExpectations(t)
	assert.Nil(t, testObj.AbstractValidator.ValidatedData)
	assert.Equal(t, map[string][]string{"field": []string{"required"}}, *testObj.AbstractValidator.Errors)
}

func TestV10ValidatorWithoutErrors(t *testing.T) {
	testObj := new(StubV10ValidatorValidated)
	structure := struct{ Field string }{Field: "some value"}

	testObj.On("Validate", structure).Return(true)
	testObj.Validate(structure)

	testObj.AssertExpectations(t)
	assert.Nil(t, testObj.AbstractValidator.Errors)
	assert.Equal(t, structure, testObj.AbstractValidator.ValidatedData)
}
