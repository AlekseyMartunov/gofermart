package handlers

import (
	"context"
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

type userDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func New(l logger, us userService) *Handler {
	return &Handler{
		logger:      l,
		userService: us,
	}
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
