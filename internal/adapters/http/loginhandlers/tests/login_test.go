package tests

import (
	"AlekseyMartunov/internal/adapters/http/loginhandlers/tests/mocks"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"AlekseyMartunov/internal/adapters/db/users/postgres"
	"AlekseyMartunov/internal/adapters/http/loginhandlers"
)

const testAPILogin = "/api/user/login"

func TestLoginHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	logger := mockloginhandlers.NewMocklogger(ctrl)
	userService := mockloginhandlers.NewMockUserService(ctrl)
	token := mockloginhandlers.NewMocktokenManager(ctrl)

	// test 1
	u1 := mockloginhandlers.UserDTO{
		Login:    "123",
		Password: "pass",
	}
	userService.EXPECT().CheckUser(ctx, u1.ToEntity()).Return("10", nil)
	token.EXPECT().CreateToken("10").Return("10", nil)
	//-----------------------------------------------------------------

	// test 2
	logger.EXPECT().Error("unmarshal error")
	//-----------------------------------------------------------------

	// test 3
	u3 := mockloginhandlers.UserDTO{
		Login:    "123",
		Password: "pass",
	}
	userService.EXPECT().CheckUser(ctx, u3.ToEntity()).Return("", postgres.ErrWrongLoginOrPassword)
	//-----------------------------------------------------------------

	// test 4
	u4 := mockloginhandlers.UserDTO{
		Login:    "123",
		Password: "pass",
	}
	userService.EXPECT().CheckUser(ctx, u4.ToEntity()).Return("", errors.New("ErrTest"))
	//-----------------------------------------------------------------

	loginHandler := loginhandlers.NewLoginHandler(logger, userService, token)

	testCase := []struct {
		name          string
		body          string
		statusCode    int
		contentType   string
		authorization string
	}{
		{
			name:          "test1",
			body:          `{"login": "123", "password": "pass"}`,
			statusCode:    http.StatusOK,
			contentType:   "application/json",
			authorization: "Bearer 10",
		},
		{
			name:          "test2",
			body:          `{"login": "123", "password": "pass"`,
			statusCode:    http.StatusBadRequest,
			contentType:   "application/json",
			authorization: "",
		},
		{
			name:          "test3",
			body:          `{"login": "123", "password": "pass"}`,
			statusCode:    http.StatusUnauthorized,
			contentType:   "application/json",
			authorization: "",
		},
		{
			name:          "test4",
			body:          `{"login": "123", "password": "pass"}`,
			statusCode:    http.StatusInternalServerError,
			contentType:   "application/json",
			authorization: "",
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
			err := loginHandler.Login(c)

			assert.Equal(t, nil, err, "Ошибка при создании запроса")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.contentType, rec.Header().Get("Content-Type"))
			assert.Equal(t, tc.authorization, rec.Header().Get("Authorization"))
		})
	}

}
