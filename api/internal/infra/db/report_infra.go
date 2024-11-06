package db

import (
	"context"
	"database/sql"
	"fmt"
	"go-temp/go-sample/api/internal/domain/model"
	"go-temp/go-sample/api/internal/domain/repo"
	"go-temp/go-sample/api/internal/pkg/database"

	"github.com/pkg/errors"
)

type reportRepository struct {
	db *database.DB
}

func NewReportRepository(db *database.DB) repo.ReportRepository {
	return &reportRepository{
		db: db,
	}
}

func (rr *reportRepository) GetReportByUserID(ctx context.Context, userID string) ([]*model.Report, error) {
	var report []*model.Report
	err := rr.db.SelectContext(ctx, &report, `
	SELECT 
		r.id,
		r.comment,
		r.thumbnail_url,
		r.user_id,
		r.recipe_id AS "recipe.id",
		r.created_at,
		r.updated_at,


		rc.id AS "recipe.id",
		rc.title AS "recipe.title",
		rc.thumbnail_url AS "recipe.thumbnail_url",
		rc.recipe AS "recipe.recipe",
		rc.category_id AS "recipe.category.id",
		rc.ingredient AS "recipe.ingredient",
		rc.created_at AS "recipe.created_at",
		rc.updated_at AS "recipe.updated_at",

		c.id AS "category.id",
		c.name AS "category.name"

	FROM 
		reports AS r
	INNER JOIN
		users AS u
	ON r.user_id = u.id
	INNER JOIN
		recipes AS rc
	ON r.recipe_id = rc.id
	INNER JOIN
		categories AS c
	ON rc.category_id = c.id
	WHERE
		r.user_id = ?
	`, userID)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no rows in infra")
		return nil, err

	}

	if err != nil {
		fmt.Println("error in infra")
		fmt.Println(err)
		return nil, err
	}

	return report, nil
}
