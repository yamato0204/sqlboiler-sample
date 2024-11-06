package usecase

import (
	"context"

	"github.com/yamato0204/sqlboiler-sample/internal/domain/model"
)

type UserUsecase interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}

type ReportUsecase interface {
	GetReportByUserID(ctx context.Context) ([]*model.Report, error)
}
