package orders

import "time"

type orderStorage interface {
	Create(number, userUUID string, time time.Time) error
	GetOwner(number string) (uuid string, err error)
}

type OrderService struct {
	orderStorage
}

func (os *OrderService) Create(number, userUUID string) error {
	return os.orderStorage.Create(number, userUUID, time.Now())
}

func (os *OrderService) GerOwner(number string) (uuid string, err error) {
	return os.orderStorage.GetOwner(number)
}
