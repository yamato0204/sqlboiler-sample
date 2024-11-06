package integration

import (
	"context"
	"go-temp/go-sample/api/internal/app/container"
	"go-temp/go-sample/api/internal/controller"
	"go-temp/go-sample/api/internal/domain/repo"
	"go-temp/go-sample/api/internal/infra/db"
	"go-temp/go-sample/api/internal/pkg/database"
	"go-temp/go-sample/api/internal/usecase"
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
