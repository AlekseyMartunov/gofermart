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
	"AlekseyMartunov/internal/adapters/http/orderhandlers"
	mock_orderhandlers "AlekseyMartunov/internal/adapters/http/orderhandlers/tests/mocks"
	"AlekseyMartunov/internal/orders"
)

const testAPICreateOrders = "/api/user/orders"

func TestCreateOrdersHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	logger := mock_orderhandlers.NewMocklogger(ctrl)
	userService := mock_orderhandlers.NewMockUserService(ctrl)
	orderService := mock_orderhandlers.NewMockOrderService(ctrl)

	orderHandler := orderhandlers.New(logger, userService, orderService)

	//test1
	o1 := mock_orderhandlers.OrderDTO{
		Number: "1234567890",
		UserID: 100,
	}
	orderService.EXPECT().Create(ctx, o1.ToEntity()).Return(nil)
	//-----------------------------------------------------------------

	//test2
	o2 := mock_orderhandlers.OrderDTO{
		Number: "1234567890",
		UserID: 100,
	}
	orderService.EXPECT().Create(ctx, o2.ToEntity()).Return(orders.ErrNotValidNumber)
	//-----------------------------------------------------------------

	//test3
	o3 := mock_orderhandlers.OrderDTO{
		Number: "1234567890",
		UserID: 100,
	}
	orderService.EXPECT().Create(ctx, o3.ToEntity()).Return(postgres.ErrOrderAlreadyCreated)
	orderService.EXPECT().GetUserID(ctx, o3.Number).Return(100, nil)
	//-----------------------------------------------------------------

	//test4
	o4 := mock_orderhandlers.OrderDTO{
		Number: "1234567890",
		UserID: 100,
	}
	orderService.EXPECT().Create(ctx, o4.ToEntity()).Return(postgres.ErrOrderAlreadyCreated)
	orderService.EXPECT().GetUserID(ctx, o3.Number).Return(-1, errors.New("ErrTest"))
	logger.EXPECT().Error("ErrTest")
	//-----------------------------------------------------------------

	//test5
	o5 := mock_orderhandlers.OrderDTO{
		Number: "1234567890",
		UserID: 100,
	}
	orderService.EXPECT().Create(ctx, o5.ToEntity()).Return(postgres.ErrOrderAlreadyCreated)
	orderService.EXPECT().GetUserID(ctx, o3.Number).Return(105, nil)
	//-----------------------------------------------------------------

	//test6
	o6 := mock_orderhandlers.OrderDTO{
		Number: "1234567890",
		UserID: 100,
	}
	orderService.EXPECT().Create(ctx, o6.ToEntity()).Return(errors.New("ErrTest"))
	logger.EXPECT().Error("ErrTest")
	//-----------------------------------------------------------------

	testCase := []struct {
		name        string
		body        string
		statusCode  int
		contentType string
		userID      int
	}{
		{
			name:        "test1",
			body:        "1234567890",
			statusCode:  http.StatusAccepted,
			contentType: "text/plain; charset=UTF-8",
			userID:      100,
		},
		{
			name:        "test2",
			body:        "1234567890",
			statusCode:  http.StatusUnprocessableEntity,
			contentType: "text/plain; charset=UTF-8",
			userID:      100,
		},
		{
			name:        "test3",
			body:        "1234567890",
			statusCode:  http.StatusOK,
			contentType: "text/plain; charset=UTF-8",
			userID:      100,
		},
		{
			name:        "test4",
			body:        "1234567890",
			statusCode:  http.StatusInternalServerError,
			contentType: "text/plain; charset=UTF-8",
			userID:      100,
		},
		{
			name:        "test5",
			body:        "1234567890",
			statusCode:  http.StatusConflict,
			contentType: "text/plain; charset=UTF-8",
			userID:      100,
		},
		{
			name:        "test6",
			body:        "1234567890",
			statusCode:  http.StatusInternalServerError,
			contentType: "text/plain; charset=UTF-8",
			userID:      100,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(
				http.MethodPost,
				testAPICreateOrders,
				strings.NewReader(tc.body))

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set("userID", tc.userID)
			err := orderHandler.SaveOrder(c)

			assert.Equal(t, nil, err, "Ошибка при создании запроса")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.contentType, rec.Header().Get("Content-Type"))

		})
	}
}
