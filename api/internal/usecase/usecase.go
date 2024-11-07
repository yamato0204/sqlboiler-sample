package usecase

import (
	"context"

	"github.com/yamato0204/sqlboiler-sample/internal/domain/entity"
	"github.com/yamato0204/sqlboiler-sample/internal/domain/model"
)

type UserUsecase interface {
	GetUserByID(ctx context.Context, userID string) (*entity.UserEntity, error)
}

type ReportUsecase interface {
	GetReportByUserID(ctx context.Context) ([]*model.Report, error)
}
