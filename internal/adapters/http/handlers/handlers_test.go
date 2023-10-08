package handlers

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
	"AlekseyMartunov/internal/adapters/http/handlers/mocks"
)

const testAPIRegister = "/api/user/register"

type testLogger struct{}

func (tl *testLogger) Info(msg string)  {}
func (tl *testLogger) Warn(msg string)  {}
func (tl *testLogger) Error(msg string) {}

func TestHandler_Register(t *testing.T) {
	l := &testLogger{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	m := mocks.NewMockUserService(ctrl)

	m.EXPECT().Register(ctx, "123", "pass").Return(nil)
	m.EXPECT().Register(ctx, "123", "passs").Return(postgres.LoginAlreadyUsedErr)
	m.EXPECT().Register(ctx, "a", "b").Return(errors.New("new err"))

	testHandler := New(l, m)

	testCase := []struct {
		name        string
		body        string
		statusCode  int
		contentType string
	}{
		{
			name:        "test1",
			body:        `{"login": "123", "password": "pass"}`,
			statusCode:  http.StatusOK,
			contentType: "application/json",
		},
		{
			name:        "test2",
			body:        `{"login": "123", "password": "passs"}`,
			statusCode:  http.StatusConflict,
			contentType: "application/json",
		},
		{
			name:        "test3",
			body:        `{"login: "123", "password": "passs"`,
			statusCode:  http.StatusBadRequest,
			contentType: "application/json",
		},
		{
			name:        "test4",
			body:        `{"login": "a", "password": "b"}`,
			statusCode:  http.StatusInternalServerError,
			contentType: "application/json",
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
			assert.Equal(t, nil, err)

			assert.Equal(t, tc.statusCode, rec.Code)
			assert.Equal(t, tc.contentType, rec.Header().Get("Content-Type"))
		})
	}
}
