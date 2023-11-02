package tests

import (
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
	"AlekseyMartunov/internal/adapters/http/loginhandlers/tests/mocks"
)

const testAPIRegister = "/api/user/register"

func TestRegisterHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	userService := mock_loginhandlers.NewMockUserService(ctrl)
	token := mock_loginhandlers.NewMocktokenManager(ctrl)
	logger := mock_loginhandlers.NewMocklogger(ctrl)

	// test 1
	u1 := mock_loginhandlers.UserDTO{
		Login:    "123",
		Password: "pass",
	}
	userService.EXPECT().Create(ctx, u1.ToEntity()).Return("10", nil)
	token.EXPECT().CreateToken("10").Return("10", nil)
	//-----------------------------------------------------------------

	//test2
	logger.EXPECT().Error("unmarshal error")
	//-----------------------------------------------------------------

	//test3
	u3 := mock_loginhandlers.UserDTO{
		Login:    "123",
		Password: "pass",
	}
	userService.EXPECT().Create(ctx, u3.ToEntity()).Return("", postgres.ErrLoginAlreadyUsed)
	//-----------------------------------------------------------------

	//test4
	u4 := mock_loginhandlers.UserDTO{
		Login:    "123",
		Password: "pass",
	}
	userService.EXPECT().Create(ctx, u4.ToEntity()).Return("", errors.New("ErrTest"))
	//-----------------------------------------------------------------

	registerHandler := loginhandlers.NewLoginHandler(logger, userService, token)

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
			body:          `{"login": "123", "password" "pass"}`,
			statusCode:    http.StatusBadRequest,
			contentType:   "application/json",
			authorization: "",
		},
		{
			name:          "test3",
			body:          `{"login": "123", "password": "pass"}`,
			statusCode:    http.StatusConflict,
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
				testAPIRegister,
				strings.NewReader(tc.body))

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := registerHandler.Register(c)

			assert.Equal(t, nil, err, "Ошибка при создании запроса")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.contentType, rec.Header().Get("Content-Type"))
			assert.Equal(t, tc.authorization, rec.Header().Get("Authorization"))
		})
	}
}
