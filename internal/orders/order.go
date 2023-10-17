package orders

import "time"

type Order struct {
	Number      string
	Status      string
	UserUUID    string
	UserID      int
	Accrual     float64
	Discount    float64
	CreatedTime time.Time
}
