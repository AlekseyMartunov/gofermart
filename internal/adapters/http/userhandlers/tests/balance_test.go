package tests

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"AlekseyMartunov/internal/adapters/http/userhandlers"
	mockuserhandlers "AlekseyMartunov/internal/adapters/http/userhandlers/tests/mocks"
	"AlekseyMartunov/internal/users"
)

const testAPIBalance = "/api/user/balance"

func TestBalanceHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	logger := mockuserhandlers.NewMocklogger(ctrl)
	userService := mockuserhandlers.NewMockUserService(ctrl)
	orderService := mockuserhandlers.NewMockOrderService(ctrl)

	//test1
	u1 := users.User{
		Bonuses:   500.5,
		Withdrawn: 42,
	}

	userService.EXPECT().Balance(ctx, 10).Return(u1, nil)
	//-----------------------------------------------------------------

	//test2
	u2 := users.User{}
	userService.EXPECT().Balance(ctx, 10).Return(u2, errors.New("ErrTest"))
	logger.EXPECT().Error("ErrTest")
	//-----------------------------------------------------------------

	userHandler := userhandlers.New(logger, userService, orderService)

	testCase := []struct {
		name        string
		body        string
		statusCode  int
		contentType string
		userID      int
	}{
		{
			name:        "test1",
			body:        fmt.Sprintf("%s\n", `{"current":500.5,"withdrawn":42}`),
			statusCode:  http.StatusOK,
			contentType: "application/json; charset=UTF-8",
			userID:      10,
		},
		{
			name:        "test2",
			body:        fmt.Sprintf("%s\n", `"Internal Server Error."`),
			statusCode:  http.StatusInternalServerError,
			contentType: "application/json; charset=UTF-8",
			userID:      10,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(
				http.MethodGet,
				testAPIBalance,
				strings.NewReader(""))

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set("userID", tc.userID)
			err := userHandler.Balance(c)

			assert.Equal(t, nil, err, "Ошибка при создании запроса")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.contentType, rec.Header().Get("Content-Type"))
			assert.Equal(t, tc.body, rec.Body.String())
		})
	}

}
