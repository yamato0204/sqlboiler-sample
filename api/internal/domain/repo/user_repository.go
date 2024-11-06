package repo

import (
	"context"
	"go-temp/go-sample/api/internal/domain/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}
