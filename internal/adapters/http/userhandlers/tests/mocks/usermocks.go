package mockuserhandlers

import (
	"AlekseyMartunov/internal/orders"
	"AlekseyMartunov/internal/users"
)

type UserDTO struct {
	Bonuses   float64 `json:"current"`
	Withdrawn float64 `json:"withdrawn"`
}

func (dto *UserDTO) FromEntity(user users.User) {
	dto.Bonuses = user.Bonuses
	dto.Withdrawn = user.Withdrawn
}

type OrderDTO struct {
	Number   string  `json:"order"`
	Discount float64 `json:"sum"`
	UserID   int
}

func (dto *OrderDTO) ToEntity() orders.Order {
	return orders.Order{
		Number:   dto.Number,
		Discount: dto.Discount,
		UserID:   dto.UserID,
	}
}
