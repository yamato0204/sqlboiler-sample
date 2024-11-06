package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/yamato0204/sqlboiler-sample/internal/controller"
)

func TestGetReportByUserID(t *testing.T) {
	ctx := context.Background()
	app, _ := Initialize(ctx)

	tests := []struct {
		name        string
		expected    []controller.Report
		expectedErr string
		preparation func()
	}{
		{
			name: "success",
			expected: []controller.Report{
				{
					ID:           "1e8b6d70-9c2f-11eb-a8b3-0242ac130003",
					Comment:      "トマトサラダのレビュー。簡単で美味しかったです！",
					ThumbnailUrl: "https://video.com/production/videos/b331637e-1067-4471-acf3-6dcc7f60d5bb/.jpg",
					UserID:       "a4e8b5d4-9c1f-11eb-a8b3-0242ac130003",
					Recipe: controller.Recipe{
						ID:           "01J6CXNF4GYJFTNTTH9TZ3MSJG",
						Title:        "トマトサラダ",
						ThumbnailUrl: "https://video.com/production/videos/b331637e-1067-4471-acf3-/.jpg",
						Recipe:       "トマトとバジルのシンプルなサラダ。",
						Category: controller.Category{
							ID:   "1",
							Name: "フレンチ",
						},
						Ingredient: `[{"amount": "100g", "ingredient": "トマト"}, {"amount": "10枚", "ingredient": "バジル"}]`,
						CreatedAt:  "2024-08-01T10:00:00Z",
						UpdatedAt:  "2024-08-01T10:00:00Z",
					},
					CreatedAt: "2024-09-01T09:00:00Z",
					UpdatedAt: "2024-09-01T09:00:00Z",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// テスト用のEchoコンテキストとレスポンスレコーダを作成
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/reports/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// 実際のハンドラを呼び出し
			err := app.ReportCtrl.GetReportByUserID(c)
			if tt.expectedErr != "" {
				assert.EqualError(t, err, tt.expectedErr)
				return
			}

			// レスポンスのステータスコードを確認
			assert.Equal(t, http.StatusOK, rec.Code)

			// レスポンスをデコード
			var actual []controller.Report
			err = json.Unmarshal(rec.Body.Bytes(), &actual)
			assert.NoError(t, err)

			// 期待される結果と一致するか確認
			assert.Equal(t, tt.expected[0].ID, actual[0].ID)
			assert.Equal(t, tt.expected[0].Comment, actual[0].Comment)
			assert.Equal(t, tt.expected[0].ThumbnailUrl, actual[0].ThumbnailUrl)
			assert.Equal(t, tt.expected[0].UserID, actual[0].UserID)
			assert.Equal(t, tt.expected[0].Recipe.ID, actual[0].Recipe.ID)
			assert.Equal(t, tt.expected[0].Recipe.Title, actual[0].Recipe.Title)
			assert.Equal(t, tt.expected[0].Recipe.ThumbnailUrl, actual[0].Recipe.ThumbnailUrl)
			assert.Equal(t, tt.expected[0].Recipe.Recipe, actual[0].Recipe.Recipe)
			assert.Equal(t, tt.expected[0].Recipe.Category.ID, actual[0].Recipe.Category.ID)
			assert.Equal(t, tt.expected[0].Recipe.Category.Name, actual[0].Recipe.Category.Name)
			assert.Equal(t, tt.expected[0].Recipe.Ingredient, actual[0].Recipe.Ingredient)
			assert.Equal(t, tt.expected[0].Recipe.CreatedAt, actual[0].Recipe.CreatedAt)
			assert.Equal(t, tt.expected[0].Recipe.UpdatedAt, actual[0].Recipe.UpdatedAt)
			assert.Equal(t, tt.expected[0].CreatedAt, actual[0].CreatedAt)
			assert.Equal(t, tt.expected[0].UpdatedAt, actual[0].UpdatedAt)
		})
	}
}
