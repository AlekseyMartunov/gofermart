package orders

import "time"

type Order struct {
	Number      string
	Status      string
	Accrual     int
	UserUUID    string
	UserID      int
	CreatedTime time.Time
}
