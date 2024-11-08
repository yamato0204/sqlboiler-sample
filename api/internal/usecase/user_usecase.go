package usecase

import (
	"context"

	"github.com/yamato0204/sqlboiler-sample/internal/domain/entity"
	"github.com/yamato0204/sqlboiler-sample/internal/domain/repo"
)

type userUsecase struct {
	repo repo.UserRepository
}

func NewUserUsecase(repo repo.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) GetUserByID(ctx context.Context, userID string) (*entity.UserEntity, error) {

	user, err := u.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
