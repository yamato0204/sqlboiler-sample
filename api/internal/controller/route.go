package controller

import (
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo,
	userCtrl *UserController,
	reportCtrl *ReportController,

) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	g := e.Group("/api")

	users := g.Group("/users")

	users.GET("/:id", userCtrl.GetUserByID)

	//reports
	reports := g.Group("/reports")

	reports.GET("/", reportCtrl.GetReportByUserID)

}
