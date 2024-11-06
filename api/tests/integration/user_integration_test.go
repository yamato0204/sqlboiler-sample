package integration

import (
	"context"
	"encoding/json"
	"go-temp/go-sample/api/internal/domain/model"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	ctx := context.Background()
	app, _ := Initialize(ctx)

	tests := []struct {
		name        string
		input       string
		expected    model.User
		expectedErr string
		preparation func()
	}{
		{
			name:  "success",
			input: "a4e8b5d4-9c1f-11eb-a8b3-0242ac130003",
			expected: model.User{
				ID:    "a4e8b5d4-9c1f-11eb-a8b3-0242ac130003",
				Name:  "Alice",
				Email: "test@test.com",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// テスト用のEchoコンテキストとレスポンスレコーダを作成
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/users/"+tt.input, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetParamNames("id")
			c.SetParamValues(tt.input)

			// 実際のハンドラを呼び出し
			err := app.UserCtrl.GetUserByID(c)
			if tt.expectedErr != "" {
				assert.EqualError(t, err, tt.expectedErr)
				return
			}

			// レスポンスのステータスコードを確認
			assert.Equal(t, http.StatusOK, rec.Code)

			// レスポンスをデコード
			var actual model.User
			err = json.Unmarshal(rec.Body.Bytes(), &actual)
			assert.NoError(t, err)
			log.Println(actual.Name)
			log.Println(actual.Email)
			log.Println(tt.expected.Name)
			log.Println(tt.expected.Email)

			// 期待される結果と一致するか確認
			assert.Equal(t, tt.expected.ID, actual.ID)
			assert.Equal(t, tt.expected.Name, actual.Name)
			assert.Equal(t, tt.expected.Email, actual.Email)
		})
	}
}
