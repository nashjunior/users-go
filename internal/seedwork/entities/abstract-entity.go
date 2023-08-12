package entities

import (
	"database/sql"
	valueobjects "school-api/internal/seedwork/value-objects"
	"time"
)

type AbstractEntityProps struct {
	uniqueEntityId *valueobjects.UniqueEntityId

	CreatedAt     time.Time `json:"createdAt"`
	UserCreatedBy *string   `json:"createdBy,omitempty"`

	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	UserUpdatedBy *string

	DeletedAt *sql.NullTime `json:"deletedAt,omitempty"`
}

type AbstractEntity struct {
	UniqueEntityId valueobjects.UniqueEntityId

	CreatedAt     time.Time `json:"createdAt"`
	UserCreatedBy *string   `json:"createdBy,omitempty"`

	UpdatedAt     *time.Time                   `json:"updatedAt,omitempty"`
	UserUpdatedBy *valueobjects.UniqueEntityId `json:"userUpdatedBy,omitempty"`

	DeletedAt *sql.NullTime `json:"deletedAt,omitempty"`
}

func NewAbstractEntity(instance AbstractEntityProps) (*AbstractEntity, error) {
	var uniqueId valueobjects.UniqueEntityId

	if instance.uniqueEntityId != nil {
		uniqueId = *instance.uniqueEntityId
	} else {
		unique, error := valueobjects.NewUniqueId(nil)

		if error != nil {
			return nil, error
		}

		uniqueId = *unique
	}

	var uniqueUserUpdate *valueobjects.UniqueEntityId

	if instance.UserUpdatedBy != nil {
		unique, error := valueobjects.NewUniqueId(instance.UserUpdatedBy)

		if error != nil {
			return nil, error
		}

		uniqueUserUpdate = unique
	}

	return &AbstractEntity{
		UniqueEntityId: uniqueId,
		CreatedAt:      instance.CreatedAt,
		UserCreatedBy:  instance.UserCreatedBy,
		UpdatedAt:      instance.UpdatedAt,
		UserUpdatedBy:  uniqueUserUpdate,
		DeletedAt:      instance.DeletedAt,
	}, nil
}
