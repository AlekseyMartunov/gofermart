package accural

import (
	"AlekseyMartunov/internal/accrualservice/requestcontroller"
	"AlekseyMartunov/internal/orders"
	"context"
	"time"
)

type logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type OrderService interface {
	Update(ctx context.Context, orders ...orders.Order) error
}

type reqController interface {
	Get(number string) (requestcontroller.OrderResponse, error)
}

type Accrual struct {
	orderService OrderService
	log          logger
}

func NewAccrual(l logger, os OrderService) *Accrual {
	return &Accrual{
		log:          l,
		orderService: os,
	}
}

func (a *Accrual) Run(ctx context.Context, ch chan orders.Order, t time.Duration) {

	go func() {
		ticker := time.NewTicker(t)
		ordersArr := make([]orders.Order, 20)
		for {
			select {
			case <-ticker.C:
				if len(ordersArr) == 0 {
					continue
				}

				a.orderService.Update(ctx, ordersArr...)

			case ord, ok := <-ch:
				if !ok {
					ch = nil
				}

				ordersArr = append(ordersArr, ord)
			}
		}
	}()
}