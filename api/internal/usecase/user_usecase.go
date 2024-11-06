package usecase

import (
	"context"
	"go-temp/go-sample/api/internal/domain/model"
	"go-temp/go-sample/api/internal/domain/repo"
)

type userUsecase struct {
	repo repo.UserRepository
}

func NewUserUsecase(repo repo.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) GetUserByID(ctx context.Context, userID string) (*model.User, error) {

	user, err := u.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
