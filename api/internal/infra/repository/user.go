package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/yamato0204/sqlboiler-sample/internal/domain/entity"
	"github.com/yamato0204/sqlboiler-sample/internal/domain/repo"
	"github.com/yamato0204/sqlboiler-sample/internal/infra/cache"
	"github.com/yamato0204/sqlboiler-sample/internal/infra/datamodel"
	"github.com/yamato0204/sqlboiler-sample/internal/infra/ifcache"
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

	cache := ifcache.ShowTable(cache.NewCache(*m))

	fmt.Println("11111111111111")
	log.Println(cache)
	fmt.Println(cache)
	fmt.Println("22222222222222")

	return entity.NewUserEntity(m), nil
}
