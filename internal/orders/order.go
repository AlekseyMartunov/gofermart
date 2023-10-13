package orders

import (
	"time"
)

type Order struct {
	Number      string
	CreatedTime time.Time
	Status      string
	UserUUID    string
	UserID      int
}
