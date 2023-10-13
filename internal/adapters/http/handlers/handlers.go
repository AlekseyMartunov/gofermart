package handlers

import (
	"AlekseyMartunov/internal/orders"
	"AlekseyMartunov/internal/users"
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

type UserService interface {
	Create(ctx context.Context, user users.User) (string, error)
	CheckUser(ctx context.Context, user users.User) (string, error)
}

type OrderService interface {
	Create(ctx context.Context, order orders.Order) error
	GetUserID(ctx context.Context, number string) (int, error)
}

type Handler struct {
	logger       logger
	userService  UserService
	orderService OrderService
}

type orderDTO struct {
	number string
	UserID int
}

func (dto *orderDTO) toEntity() orders.Order {
	return orders.Order{
		Number: dto.number,
		UserID: dto.UserID,
	}
}

func New(l logger, us UserService, os OrderService) *Handler {
	return &Handler{
		logger:       l,
		userService:  us,
		orderService: os,
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
