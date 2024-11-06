package usecase

import (
	"context"
	"fmt"

	"github.com/yamato0204/sqlboiler-sample/internal/domain/model"
	"github.com/yamato0204/sqlboiler-sample/internal/domain/repo"
)

type reportUsecase struct {
	repo repo.ReportRepository
}

func NewReportUsecase(repo repo.ReportRepository) ReportUsecase {
	return &reportUsecase{
		repo: repo,
	}
}

func (r *reportUsecase) GetReportByUserID(ctx context.Context) ([]*model.Report, error) {

	userID := "a4e8b5d4-9c1f-11eb-a8b3-0242ac130003"

	report, err := r.repo.GetReportByUserID(ctx, userID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error in usecase")
		return nil, err
	}

	return report, nil
}
