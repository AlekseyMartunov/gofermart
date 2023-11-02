package tests

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"AlekseyMartunov/internal/adapters/db/users/postgres"
	"AlekseyMartunov/internal/adapters/http/userhandlers"
	mock_userhandlers "AlekseyMartunov/internal/adapters/http/userhandlers/tests/mocks"
	"AlekseyMartunov/internal/users"
)

func TestWithdrawalsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	logger := mock_userhandlers.NewMocklogger(ctrl)
	userService := mock_userhandlers.NewMockUserService(ctrl)
	orderService := mock_userhandlers.NewMockOrderService(ctrl)

	//test1
	t1, _ := time.Parse(time.RFC3339, "2020-12-09T16:09:57+03:00")
	t2, _ := time.Parse(time.RFC3339, "2020-12-09T16:10:57+03:00")

	arr1 := []users.HistoryElement{
		{
			Order:        "648474",
			Amount:       7485.9,
			WriteOffTime: t1,
		},
		{
			Order:        "648474",
			Amount:       7485.9,
			WriteOffTime: t2,
		},
	}
	userService.EXPECT().GetHistory(ctx, 100).Return(arr1, nil)
	//-----------------------------------------------------------------

	//test2
	userService.EXPECT().GetHistory(ctx, 100).Return(nil, errors.New("ErrTest"))
	logger.EXPECT().Error("ErrTest")
	//-----------------------------------------------------------------

	//test3
	userService.EXPECT().GetHistory(ctx, 100).Return(nil, postgres.ErrEmptyHistory)
	//-----------------------------------------------------------------

	userHandler := userhandlers.New(logger, userService, orderService)

	body1 := fmt.Sprintf("%s\n", `[{"order":"648474","sum":7485.9,"processed_at":"2020-12-09T16:09:57+03:00"},{"order":"648474","sum":7485.9,"processed_at":"2020-12-09T16:10:57+03:00"}]`)

	testCase := []struct {
		name        string
		body        string
		statusCode  int
		contentType string
		userID      int
	}{
		{
			name:        "test1",
			body:        body1,
			statusCode:  http.StatusOK,
			contentType: "application/json; charset=UTF-8",
			userID:      100,
		},
		{
			name:        "test2",
			body:        fmt.Sprintf("%s\n", `"Internal Server Error."`),
			statusCode:  http.StatusInternalServerError,
			contentType: "application/json; charset=UTF-8",
			userID:      100,
		},
		{
			name:        "test3",
			body:        fmt.Sprintf("%s\n", `"empty history"`),
			statusCode:  http.StatusNoContent,
			contentType: "application/json; charset=UTF-8",
			userID:      100,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(
				http.MethodGet,
				testAPIWithdraw,
				strings.NewReader(""))

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set("userID", tc.userID)
			err := userHandler.Withdrawals(c)

			assert.Equal(t, nil, err, "Ошибка при создании запроса")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.contentType, rec.Header().Get("Content-Type"))
			assert.Equal(t, tc.body, rec.Body.String())

		})
	}
}
