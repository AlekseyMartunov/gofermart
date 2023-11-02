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

	"AlekseyMartunov/internal/adapters/db/orders/postgres"
	"AlekseyMartunov/internal/adapters/http/userhandlers"
	mock_userhandlers "AlekseyMartunov/internal/adapters/http/userhandlers/tests/mocks"
	"AlekseyMartunov/internal/orders"
)

const testAPIWithdraw = "/api/user/balance/withdraw"

func TestWithdrawHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	logger := mock_userhandlers.NewMocklogger(ctrl)
	userService := mock_userhandlers.NewMockUserService(ctrl)
	orderService := mock_userhandlers.NewMockOrderService(ctrl)

	//test1
	o1 := mock_userhandlers.OrderDTO{
		Number:   "2377225624",
		Discount: 751.0,
		UserID:   100,
	}
	orderService.EXPECT().AddDiscount(ctx, o1.ToEntity()).Return(nil)
	//-----------------------------------------------------------------

	//test2
	logger.EXPECT().Error("Unmarshal error")
	//-----------------------------------------------------------------

	//test3
	o3 := mock_userhandlers.OrderDTO{
		Number:   "2377225624",
		Discount: 751.0,
		UserID:   100,
	}
	orderService.EXPECT().AddDiscount(ctx, o3.ToEntity()).Return(orders.ErrNotValidNumber)
	//-----------------------------------------------------------------

	//test4
	o4 := mock_userhandlers.OrderDTO{
		Number:   "2377225624",
		Discount: 751.0,
		UserID:   100,
	}
	orderService.EXPECT().AddDiscount(ctx, o4.ToEntity()).Return(postgres.ErrNotEnoughMoney)
	//-----------------------------------------------------------------

	//test5
	o5 := mock_userhandlers.OrderDTO{
		Number:   "2377225624",
		Discount: 751.0,
		UserID:   100,
	}
	orderService.EXPECT().AddDiscount(ctx, o5.ToEntity()).Return(errors.New("ErrTest"))
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
			body:        `{"order": "2377225624", "sum": 751}`,
			statusCode:  http.StatusOK,
			contentType: "application/json; charset=UTF-8",
			userID:      100,
		},
		{
			name:        "test2",
			body:        `{"order": "2377225624 "sum": 751}`,
			statusCode:  http.StatusBadRequest,
			contentType: "application/json; charset=UTF-8",
			userID:      100,
		},
		{
			name:        "test3",
			body:        `{"order": "2377225624", "sum": 751}`,
			statusCode:  http.StatusUnprocessableEntity,
			contentType: "application/json; charset=UTF-8",
			userID:      100,
		},
		{
			name:        "test4",
			body:        `{"order": "2377225624", "sum": 751}`,
			statusCode:  http.StatusPaymentRequired,
			contentType: "application/json; charset=UTF-8",
			userID:      100,
		},
		{
			name:        "test5",
			body:        `{"order": "2377225624", "sum": 751}`,
			statusCode:  http.StatusInternalServerError,
			contentType: "application/json; charset=UTF-8",
			userID:      100,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(
				http.MethodPost,
				testAPIWithdraw,
				strings.NewReader(tc.body))

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set("userID", tc.userID)
			err := userHandler.Withdraw(c)

			assert.Equal(t, nil, err, "Ошибка при создании запроса")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.contentType, rec.Header().Get("Content-Type"))

		})
	}
}
