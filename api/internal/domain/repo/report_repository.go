package repo

import (
	"context"

	"github.com/yamato0204/sqlboiler-sample/internal/domain/model"
)

type ReportRepository interface {
	GetReportByUserID(ctx context.Context, userID string) ([]*model.Report, error)
}
