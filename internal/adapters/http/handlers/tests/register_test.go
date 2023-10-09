package tests

import (
	"AlekseyMartunov/internal/adapters/http/handlers"
	"AlekseyMartunov/internal/utils/tokenmanager"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"AlekseyMartunov/internal/adapters/db/users/postgres"
	"AlekseyMartunov/internal/adapters/http/handlers/mocks"
)

const testAPIRegister = "/api/user/register"

func TestHandler_Register(t *testing.T) {
	l := &testLogger{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	m := mocks.NewMockUserService(ctrl)

	m.EXPECT().Create(ctx, "123", "pass").Return("78457-34394-394839", nil)
	m.EXPECT().Create(ctx, "123", "passs").Return("", postgres.ErrLoginAlreadyUsed)
	m.EXPECT().Create(ctx, "a", "b").Return("", errors.New("new err"))

	tk := tokenmanager.New(time.Hour, []byte("key"))

	testHandler := handlers.New(l, m, tk)

	testCase := []struct {
		name        string
		body        string
		statusCode  int
		contentType string
		checkToken  bool
	}{
		{
			name:        "test1",
			body:        `{"login": "123", "password": "pass"}`,
			statusCode:  http.StatusOK,
			contentType: "application/json",
			checkToken:  true,
		},
		{
			name:        "test2",
			body:        `{"login": "123", "password": "passs"}`,
			statusCode:  http.StatusConflict,
			contentType: "application/json",
			checkToken:  false,
		},
		{
			name:        "test3",
			body:        `{"login: "123", "password": "passs"`,
			statusCode:  http.StatusBadRequest,
			contentType: "application/json",
			checkToken:  false,
		},
		{
			name:        "test4",
			body:        `{"login": "a", "password": "b"}`,
			statusCode:  http.StatusInternalServerError,
			contentType: "application/json",
			checkToken:  false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

			e := echo.New()

			req := httptest.NewRequest(
				http.MethodPost,
				testAPIRegister,
				strings.NewReader(tc.body))

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			err := testHandler.Register(c)
			assert.NoError(t, err, "Хендлер вернул ошибку")

			assert.Equal(t, tc.statusCode, rec.Code,
				"Не совпадает статус код")

			assert.Equal(t, tc.contentType, rec.Header().Get("Content-Type"),
				"Не совпадает Content-Type")
			if tc.checkToken {
				assert.NotEmpty(t, rec.Header().Get("Authorization"),
					"Отсутствует токен в заголовке")
			}
		})
	}
}
