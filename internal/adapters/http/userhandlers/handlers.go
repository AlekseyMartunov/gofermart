package userhandlers

import (
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

type UserHandlers struct {
	logger      logger
	userService UserService
}

func New(l logger, us UserService) *UserHandlers {
	return &UserHandlers{
		logger:      l,
		userService: us,
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

func (h *UserHandlers) Withdraw(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (h *UserHandlers) Withdrawals(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
