package repo

import (
	"context"
	"go-temp/go-sample/api/internal/domain/model"
)

type ReportRepository interface {
	GetReportByUserID(ctx context.Context, userID string) ([]*model.Report, error)
}
