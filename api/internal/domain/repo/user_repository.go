package repo

import (
	"context"

	"github.com/yamato0204/sqlboiler-sample/internal/domain/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}
