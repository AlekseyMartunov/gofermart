package userhandlers

import (
	"AlekseyMartunov/internal/orders"
	"AlekseyMartunov/internal/users"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type UserService interface {
	Balance(ctx context.Context, userID int) (users.User, error)
}

type OrderService interface {
	AddDiscount(ctx context.Context, order orders.Order) error
}

type UserHandlers struct {
	logger       logger
	userService  UserService
	orderService OrderService
}

func New(l logger, us UserService, os OrderService) *UserHandlers {
	return &UserHandlers{
		logger:       l,
		userService:  us,
		orderService: os,
	}
}

type userDTO struct {
	Bonuses   float64 `json:"current"`
	Withdrawn float64 `json:"withdrawn"`
}

func (dto *userDTO) fromEntity(user users.User) {
	dto.Bonuses = user.Bonuses
	dto.Withdrawn = user.Withdrawn
}

type orderDTO struct {
	Number   string  `json:"order"`
	Discount float64 `json:"sum"`
	UserID   int
}

func (dto *orderDTO) toEntity() orders.Order {
	return orders.Order{
		Number:   dto.Number,
		Discount: dto.Discount,
		UserID:   dto.UserID,
	}
}

func (h *UserHandlers) Withdrawals(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
