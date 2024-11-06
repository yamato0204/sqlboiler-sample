package integration

import (
	"context"
	"go-temp/go-sample/api/internal/app/container"
	"go-temp/go-sample/api/internal/pkg/database"
	"log"
)

func Initialize(ctx context.Context) (*container.Container, InfraInstance) {

	dbConfig := NewDBConfig()
	databaseDBConfig := ConvertDBEnv(dbConfig)
	db := database.NewDB(databaseDBConfig)

	app, infraInstances := NewTestController(ctx, db)
	err := db.Migrate()
	if err != nil {
		log.Fatalf("failed to migrate: %v\n", err)
		return nil, infraInstances
	}
	log.Println("migrated")

	return app, infraInstances
}

func ConvertDBEnv(cfg *DBConfig) *database.DBConfig {
	return &database.DBConfig{
		User:     cfg.User,
		DBName:   cfg.DBName,
		RootPass: cfg.RootPass,
		Host:     cfg.Host,
		Port:     cfg.Port,
	}
}
