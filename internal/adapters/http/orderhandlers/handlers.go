package orderhandlers

import (
	"AlekseyMartunov/internal/orders"
	"AlekseyMartunov/internal/users"
	"context"
	"time"
)

//go:generate mockgen -source handlers.go  -destination tests/mocks/mock_orderhandlers.go

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
	GetOrders(ctx context.Context, userID int) ([]orders.Order, error)
}

type OrderHandler struct {
	logger       logger
	userService  UserService
	orderService OrderService
}

func New(l logger, us UserService, os OrderService) *OrderHandler {
	return &OrderHandler{
		logger:       l,
		userService:  us,
		orderService: os,
	}
}

type orderDTO struct {
	Number      string    `json:"number"`
	Accrual     float64   `json:"accrual,omitempty"`
	CreatedTime time.Time `json:"uploaded_at"`
	Status      string    `json:"status"`
	UserID      int       `json:"-"`
}

func fromEntity(orders []orders.Order) []orderDTO {
	resDTO := make([]orderDTO, len(orders))

	for i, o := range orders {
		resDTO[i].Number = o.Number
		resDTO[i].Accrual = o.Accrual
		resDTO[i].CreatedTime = o.CreatedTime
		resDTO[i].Status = o.Status
	}
	return resDTO
}

func (dto *orderDTO) toEntity() orders.Order {
	return orders.Order{
		Number: dto.Number,
		UserID: dto.UserID,
	}
}
