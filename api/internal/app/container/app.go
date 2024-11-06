package container

import (
	"errors"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/yamato0204/sqlboiler-sample/internal/controller"
	"github.com/yamato0204/sqlboiler-sample/internal/pkg/database"
)

type Container struct {
	UserCtrl   *controller.UserController
	ReportCtrl *controller.ReportController
}

func NewCtrl(
	userCtrl *controller.UserController,
	reportCtrl *controller.ReportController,

) *Container {
	return &Container{
		UserCtrl:   userCtrl,
		ReportCtrl: reportCtrl,
	}
}

type App struct {
	e  *echo.Echo
	db *database.DB
}

func NewApp(db *database.DB, container *Container) (*App, error) {

	e := echo.New()

	//	e.Validator = NewValidator()
	//
	controller.SetupRoutes(e, container.UserCtrl, container.ReportCtrl)

	return &App{
		e:  e,
		db: db,
	}, nil
}

func (a *App) Migrate() error {
	return a.db.Migrate()
}

func (a *App) Start() error {
	return a.e.Start(":8080")
}

func (a *App) Close() error {
	return errors.Join(
		a.db.Close(),
	)
}

type echovalidator struct {
	validator *validator.Validate
}

func NewValidator() *echovalidator {
	return &echovalidator{
		validator: validator.New(),
	}
}

func (ev *echovalidator) Validate(i interface{}) error {
	return ev.validator.Struct(i)
}
