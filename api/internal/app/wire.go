//go:build wireinject

//go:generate wire $GOFILE

package app

import (
	"context"
	"go-temp/go-sample/api/internal/app/config"
	"go-temp/go-sample/api/internal/app/container"
	"go-temp/go-sample/api/internal/controller"
	"go-temp/go-sample/api/internal/infra/db"
	"go-temp/go-sample/api/internal/pkg/database"
	"go-temp/go-sample/api/internal/usecase"

	"github.com/google/wire"
)

func New(context.Context) (*container.App, error) {
	wire.Build(

		container.NewCtrl,
		container.NewApp,

		// DB
		database.NewDB,
		ConvertDBEnv,
		config.NewDBConfig,

		controller.NewUserController,
		controller.NewReportController,

		usecase.NewUserUsecase,
		usecase.NewReportUsecase,

		db.NewUserRepository,
		db.NewReportRepository,
	)

	return nil, nil
}

func ConvertDBEnv(cfg *config.DBConfig) *database.DBConfig {
	return &database.DBConfig{
		User:     cfg.User,
		DBName:   cfg.DBName,
		RootPass: cfg.RootPass,
		Host:     cfg.Host,
		Port:     cfg.Port,
	}
}
