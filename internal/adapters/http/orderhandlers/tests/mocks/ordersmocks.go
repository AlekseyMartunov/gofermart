package mock_orderhandlers

import (
	"time"

	"AlekseyMartunov/internal/orders"
)

type OrderDTO struct {
	Number      string    `json:"number"`
	Accrual     float64   `json:"accrual,omitempty"`
	CreatedTime time.Time `json:"uploaded_at"`
	Status      string    `json:"status"`
	UserID      int       `json:"-"`
}

func FromEntity(orders []orders.Order) []OrderDTO {
	resDTO := make([]OrderDTO, len(orders))

	for i, o := range orders {
		resDTO[i].Number = o.Number
		resDTO[i].Accrual = o.Accrual
		resDTO[i].CreatedTime = o.CreatedTime
		resDTO[i].Status = o.Status
	}
	return resDTO
}

func (dto *OrderDTO) ToEntity() orders.Order {
	return orders.Order{
		Number: dto.Number,
		UserID: dto.UserID,
	}
}
