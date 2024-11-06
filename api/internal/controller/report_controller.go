package controller

import (
	"fmt"
	"go-temp/go-sample/api/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Report struct {
	ID           string `json:"id"`
	Comment      string `json:"comment"`
	ThumbnailUrl string `json:"thumbnail_url"`
	UserID       string `json:"user_id"`
	Recipe       Recipe `json:"recipe_id"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Recipe struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	ThumbnailUrl string   `json:"thumbnail_url"`
	Recipe       string   `json:"recipe"`
	Category     Category `json:"category"`
	Ingredient   string   `json:"ingredient"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ReportController struct {
	rc usecase.ReportUsecase
}

func NewReportController(rc usecase.ReportUsecase) *ReportController {
	return &ReportController{
		rc: rc,
	}
}

func (ct *ReportController) GetReportByUserID(c echo.Context) error {
	reports, err := ct.rc.GetReportByUserID(c.Request().Context())
	if err != nil {
		fmt.Println(err)
		fmt.Println("error in controller")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//Create a slice of response structs to hold the report data
	var reportResponses []Report
	for _, report := range reports {
		reportResponses = append(reportResponses, Report{
			ID:           report.ID,
			Comment:      report.Comment,
			ThumbnailUrl: report.ThumbnailUrl,
			UserID:       report.UserID,
			Recipe: Recipe{
				ID:           report.Recipe.ID,
				Title:        report.Recipe.Title,
				ThumbnailUrl: report.Recipe.ThumbnailUrl,
				Recipe:       report.Recipe.Recipe,
				Category: Category{
					ID:   report.Category.ID,
					Name: report.Category.Name,
				},
				Ingredient: report.Recipe.Ingredient,
				CreatedAt:  report.Recipe.CreatedAt,
				UpdatedAt:  report.Recipe.UpdatedAt,
			},
			CreatedAt: report.CreatedAt,
			UpdatedAt: report.UpdatedAt,
		})
	}

	// Return the list of reports as a JSON response
	return c.JSON(http.StatusOK, reportResponses)
}
