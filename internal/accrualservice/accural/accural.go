package accural

import (
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

type Accrual struct {
	orderService   OrderService
	log            logger
	ordersToUpdate []orders.Order
}

func NewAccrual(l logger, os OrderService) *Accrual {
	ordersArr := make([]orders.Order, 20)

	return &Accrual{
		log:            l,
		orderService:   os,
		ordersToUpdate: ordersArr,
	}
}

func (a *Accrual) Run(ctx context.Context, ch chan orders.Order, t time.Duration) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case ord, ok := <-ch:
				if !ok {
					return
				}
				a.ordersToUpdate = append(a.ordersToUpdate, ord)
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(t)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				a.orderService.Update(ctx, a.ordersToUpdate...)
				a.ordersToUpdate = a.ordersToUpdate[:]
			}
		}
	}()
}
