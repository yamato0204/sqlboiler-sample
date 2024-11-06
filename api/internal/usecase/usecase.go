package usecase

import (
	"context"
	"go-temp/go-sample/api/internal/domain/model"
)

type UserUsecase interface {
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
}

type ReportUsecase interface {
	GetReportByUserID(ctx context.Context) ([]*model.Report, error)
}
