package integration

import (
	"context"

	"github.com/yamato0204/sqlboiler-sample/internal/app/container"
	"github.com/yamato0204/sqlboiler-sample/internal/controller"
	"github.com/yamato0204/sqlboiler-sample/internal/domain/repo"
	"github.com/yamato0204/sqlboiler-sample/internal/infra/db"
	"github.com/yamato0204/sqlboiler-sample/internal/pkg/database"
	"github.com/yamato0204/sqlboiler-sample/internal/usecase"
)

type InfraInstance struct {
	UserRepository   repo.UserRepository
	ReportRepository repo.ReportRepository
}

func NewTestController(ctx context.Context, database *database.DB) (*container.Container, InfraInstance) {

	ur := db.NewUserRepository(database)
	rr := db.NewReportRepository(database)

	uu := usecase.NewUserUsecase(ur)
	ru := usecase.NewReportUsecase(rr)

	userController := controller.NewUserController(uu)
	reportController := controller.NewReportController(ru)
	controller := container.NewCtrl(userController, reportController)

	InfraInstance := InfraInstance{
		UserRepository:   ur,
		ReportRepository: rr,
	}

	return controller, InfraInstance

}
