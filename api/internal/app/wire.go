//go:build wireinject

//go:generate wire $GOFILE

package app

import (
	"context"

	"github.com/yamato0204/sqlboiler-sample/internal/app/config"
	"github.com/yamato0204/sqlboiler-sample/internal/app/container"
	"github.com/yamato0204/sqlboiler-sample/internal/controller"
	"github.com/yamato0204/sqlboiler-sample/internal/infra/db"
	"github.com/yamato0204/sqlboiler-sample/internal/infra/repository"
	"github.com/yamato0204/sqlboiler-sample/internal/pkg/database"
	"github.com/yamato0204/sqlboiler-sample/internal/usecase"

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

		repository.NewUserRepository,
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
