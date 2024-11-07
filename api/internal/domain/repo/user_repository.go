package repo

import (
	"context"

	"github.com/yamato0204/sqlboiler-sample/internal/domain/entity"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.UserEntity, error)
}
