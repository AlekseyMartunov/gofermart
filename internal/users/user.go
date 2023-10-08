package users

import "AlekseyMartunov/internal/orders"

type User struct {
	ID       string
	UUID     string
	Login    string
	Password string
	Money    int
	Balls    int
	orders   []orders.Order
}
