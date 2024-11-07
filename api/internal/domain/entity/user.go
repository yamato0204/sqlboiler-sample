package entity

import (
	"time"

	"github.com/yamato0204/sqlboiler-sample/internal/infra/datamodel"
)

type UserEntity struct {
	id        string
	name      string
	email     string
	createdAt time.Time
	updatedAt time.Time
}

func NewUserEntity(m *datamodel.User) *UserEntity {
	createdAt := time.Time{}
	if m.CreatedAt.Valid {
		createdAt = m.CreatedAt.Time
	}

	updatedAt := time.Time{}
	if m.UpdatedAt.Valid {
		updatedAt = m.UpdatedAt.Time
	}

	return &UserEntity{
		id:        m.ID,
		name:      m.Name,
		email:     m.Email,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (e *UserEntity) UserID() string {
	return e.id
}

func (e *UserEntity) Name() string {
	return e.name
}

func (e *UserEntity) Email() string {
	return e.email
}
