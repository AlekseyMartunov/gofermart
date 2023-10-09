package orders

import (
	"AlekseyMartunov/internal/users"
	"time"
)

type Order struct {
	number      string
	createdTime time.Time
	status      string
	user        users.User
}
