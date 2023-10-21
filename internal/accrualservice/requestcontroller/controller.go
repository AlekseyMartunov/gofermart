package requestcontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"

	"AlekseyMartunov/internal/orders"
)

type OrderResponse struct {
	Number  string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
}

type logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type RequestAccrual struct {
	host   string
	url    string
	client *resty.Client
	log    logger
}

func New(host, url string, l logger) *RequestAccrual {
	r := resty.New()
	return &RequestAccrual{
		host:   host,
		url:    url,
		client: r,
		log:    l,
	}
}

func (r *RequestAccrual) Run(ctx context.Context, ch chan string) chan orders.Order {
	orderChan := make(chan orders.Order)
	go func() {
		defer close(orderChan)
		for {
			select {
			case <-ctx.Done():
				return

			case number, ok := <-ch:
				if !ok {
					ch = nil
				}

				order := r.get(number)
				orderChan <- order
			}
		}
	}()
	return orderChan
}

func (r *RequestAccrual) get(number string) orders.Order {
	o := OrderResponse{}
	order := orders.Order{}

	client := resty.New()

	resp, err := client.R().
		Get(r.host + r.url + number)

	r.log.Info(fmt.Sprintf("url:", r.host+r.url+number))

	err = json.Unmarshal(resp.Body(), &o)
	if err != nil {
		r.log.Error(err.Error())
		return order
	}

	if err != nil {
		r.log.Error(err.Error())
		return order
	}

	order.Number = o.Number
	order.Status = o.Status
	order.Accrual = o.Accrual

	r.log.Info(fmt.Sprintf("order: %s, %s, %d", order.Number, order.Status, order.Accrual))

	return order
}
