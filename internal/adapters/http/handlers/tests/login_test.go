package tests

import (
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
	"AlekseyMartunov/internal/adapters/http/handlers"
	"AlekseyMartunov/internal/adapters/http/handlers/mocks"
	"AlekseyMartunov/internal/utils/tokenmanager"
)

const testAPILogin = "/api/user/login"

func TestHandler_Login(t *testing.T) {
	l := &testLogger{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	m := mocks.NewMockUserService(ctrl)

	m.EXPECT().CheckUserUUID(ctx, "login", "pass").Return("userUUID", nil)
	m.EXPECT().CheckUserUUID(ctx, "wrongLogin", "pass").Return("", postgres.WrongLoginOrPasswordErr)
	m.EXPECT().CheckUserUUID(ctx, "login", "password").Return("", errors.New("another err"))

	tk := tokenmanager.New(time.Hour, []byte("key"))

	testHandler := handlers.New(l, m, tk)

	testCase := []struct {
		name        string
		body        string
		checkToken  bool
		statusCode  int
		contentType string
	}{
		{
			name:        "test1",
			body:        `{"login": "login", "password": "pass"}`,
			checkToken:  true,
			statusCode:  http.StatusOK,
			contentType: "application/json",
		},
		{
			name:        "test2",
			body:        `"login": "login", "password": "pass"}`,
			checkToken:  false,
			statusCode:  http.StatusBadRequest,
			contentType: "application/json",
		},
		{
			name:        "test3",
			body:        `{"login": "wrongLogin", "password": "pass"}`,
			checkToken:  false,
			statusCode:  http.StatusUnauthorized,
			contentType: "application/json",
		},
		{
			name:        "test4",
			body:        `{"login": "login", "password": "password"}`,
			checkToken:  false,
			statusCode:  http.StatusInternalServerError,
			contentType: "application/json",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

			e := echo.New()

			req := httptest.NewRequest(
				http.MethodPost,
				testAPILogin,
				strings.NewReader(tc.body))

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			err := testHandler.Login(c)

			assert.NoError(t, err, "Хендлер вернул ошибку")

			assert.Equal(t, tc.statusCode, rec.Code,
				"Не совпадает статус код")

			assert.Equal(t, tc.contentType, rec.Header().Get("Content-Type"),
				"Не совпадает Content-Type")

			if tc.checkToken {
				assert.NotEmpty(t, rec.Header().Get("Authorization"),
					"В заголовке отсутсвует токен")
			}
		})
	}
}
