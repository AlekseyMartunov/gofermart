package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:generate mockgen -package mocks -destination mocks/mock_userservice.go . UserService

type logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type tokenManager interface {
	CreateToken(uuid string) (string, error)
}

type userService interface {
	Create(ctx context.Context, login, password string) error
	CheckUser(ctx context.Context, login, password string) (string, error)
}

type Handler struct {
	logger
	userService
	tokenManager
}

type userDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func New(l logger, us userService, tk tokenManager) *Handler {
	return &Handler{
		logger:       l,
		userService:  us,
		tokenManager: tk,
	}
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
