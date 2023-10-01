package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type userService interface {
	Register(ctx context.Context, login, password string) error
}

type Handler struct {
	logger
	userService
}

func New(l logger, us userService) *Handler {
	return &Handler{
		logger:      l,
		userService: us,
	}
}

func (h *Handler) Register(c echo.Context) error {
	defer c.Request().Body.Close()

	b, err := io.ReadAll(c.Request().Body)

	h.userService.Register(context.Background(), "", "")
	return c.String(http.StatusOK, "OK")
}

func (h *Handler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *Handler) GetOrders(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *Handler) SaveOrder(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *Handler) Balance(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *Handler) Withdraw(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *Handler) Withdrawals(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
