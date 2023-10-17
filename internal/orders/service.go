package orders

import (
	"context"
	"errors"
	"strconv"
	"time"

	"AlekseyMartunov/internal/utils/luhnalgorithm"
)

var ErrNotValidNumber = errors.New("not valid order number")

type orderStorage interface {
	Create(ctx context.Context, order Order) error
	GetUserID(ctx context.Context, number string) (int, error)
	GetOrders(ctx context.Context, userID int) ([]Order, error)
	AddDiscount(ctx context.Context, order Order) error
}

type OrderService struct {
	repo orderStorage
}

func NewOrderService(s orderStorage) *OrderService {
	return &OrderService{
		repo: s,
	}
}

func (os *OrderService) Create(ctx context.Context, order Order) error {
	intNumber, err := strconv.Atoi(order.Number)
	if err != nil {
		return ErrNotValidNumber
	}

	if !luhnalgorithm.IsValid(intNumber) {
		return ErrNotValidNumber
	}

	order.CreatedTime = time.Now()
	return os.repo.Create(ctx, order)
}

func (os *OrderService) AddDiscount(ctx context.Context, order Order) error {
	intNumber, err := strconv.Atoi(order.Number)
	if err != nil {
		return ErrNotValidNumber
	}

	if !luhnalgorithm.IsValid(intNumber) {
		return ErrNotValidNumber
	}

	order.CreatedTime = time.Now()

	return os.repo.AddDiscount(ctx, order)
}

func (os *OrderService) GetUserID(ctx context.Context, number string) (int, error) {
	return os.repo.GetUserID(ctx, number)
}

func (os *OrderService) GetOrders(ctx context.Context, userID int) ([]Order, error) {
	return os.repo.GetOrders(ctx, userID)
}
