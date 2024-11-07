package repository

import (
	"context"

	"github.com/yamato0204/sqlboiler-sample/internal/domain/entity"
	"github.com/yamato0204/sqlboiler-sample/internal/domain/repo"
	"github.com/yamato0204/sqlboiler-sample/internal/infra/datamodel"
	"github.com/yamato0204/sqlboiler-sample/internal/pkg/database"
)

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) repo.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) GetUserByID(ctx context.Context, userID string) (*entity.UserEntity, error) {
	m, err := datamodel.Users(
		datamodel.UserWhere.ID.EQ(userID),
	).One(ctx, ur.db.DB)
	if err != nil {
		return nil, err
	}

	return entity.NewUserEntity(m), nil
}
