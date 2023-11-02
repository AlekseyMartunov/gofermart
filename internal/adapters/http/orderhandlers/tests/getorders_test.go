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

	"AlekseyMartunov/internal/adapters/db/orders/postgres"
	"AlekseyMartunov/internal/adapters/http/orderhandlers"
	mock_orderhandlers "AlekseyMartunov/internal/adapters/http/orderhandlers/tests/mocks"
	"AlekseyMartunov/internal/orders"
)

const testAPIGetOrders = "/api/user/orders"

func TestGetOrdersHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	logger := mock_orderhandlers.NewMocklogger(ctrl)
	userService := mock_orderhandlers.NewMockUserService(ctrl)
	orderService := mock_orderhandlers.NewMockOrderService(ctrl)

	orderHandler := orderhandlers.New(logger, userService, orderService)

	//test1
	t1, _ := time.Parse(time.RFC3339, "2020-12-09T16:09:57+03:00")
	t2, _ := time.Parse(time.RFC3339, "2020-12-09T16:10:57+03:00")

	arr1 := []orders.Order{
		{
			Number:      "12345",
			Status:      "PROCESSED",
			Accrual:     584.95,
			CreatedTime: t1,
		},
		{
			Number:      "74833",
			Status:      "INVALID",
			Accrual:     743.3332,
			CreatedTime: t2,
		},
	}

	orderService.EXPECT().GetOrders(ctx, 100).Return(arr1, nil)
	//-----------------------------------------------------------------

	//test2
	orderService.EXPECT().GetOrders(ctx, 100).Return(nil, postgres.ErrEmptyResult)
	//-----------------------------------------------------------------

	//test3
	orderService.EXPECT().GetOrders(ctx, 100).Return(nil, errors.New("ErrTest"))
	logger.EXPECT().Error("ErrTest")
	//-----------------------------------------------------------------

	body1 := fmt.Sprintf("%s\n", "[{\"number\":\"12345\",\"accrual\":584.95,\"uploaded_at\":\"2020-12-09T16:09:57+03:00\",\"status\":\"PROCESSED\"},{\"number\":\"74833\",\"accrual\":743.3332,\"uploaded_at\":\"2020-12-09T16:10:57+03:00\",\"status\":\"INVALID\"}]")
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
			body:        fmt.Sprintf("%s\n", `"You do not have orders"`),
			statusCode:  http.StatusNoContent,
			contentType: "application/json; charset=UTF-8",
			userID:      100,
		},
		{
			name:        "test3",
			body:        fmt.Sprintf("%s\n", `"Internal Server Error."`),
			statusCode:  http.StatusInternalServerError,
			contentType: "application/json; charset=UTF-8",
			userID:      100,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			req := httptest.NewRequest(
				http.MethodGet,
				testAPIGetOrders,
				strings.NewReader(""))

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set("userID", tc.userID)
			err := orderHandler.GetOrders(c)

			assert.Equal(t, nil, err, "Ошибка при создании запроса")

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.contentType, rec.Header().Get("Content-Type"))
			assert.Equal(t, tc.body, rec.Body.String())

		})
	}
}
