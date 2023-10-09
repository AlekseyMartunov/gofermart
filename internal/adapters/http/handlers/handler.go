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

type UserService interface {
	Create(ctx context.Context, login, password string) (string, error)
	CheckUser(ctx context.Context, login, password string) (string, error)
}

type OrderService interface {
	Create(number, userUUID string) error
	GetOwner(number string) (uuid string, err error)
}

type Handler struct {
	logger
	userService  UserService
	orderService OrderService
	tokenManager
}

type userDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type orderDTO struct {
	number string
}

func New(l logger, us UserService, tk tokenManager) *Handler {
	return &Handler{
		logger:       l,
		userService:  us,
		tokenManager: tk,
	}
}

func (h *Handler) GetOrders(c echo.Context) error {
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
