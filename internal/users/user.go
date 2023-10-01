package users

import "AlekseyMartunov/internal/orders"

type User struct {
	Login    string
	Password string
	Money    int
	Balls    int
	orders   []orders.Order
}
