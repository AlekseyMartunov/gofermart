package requestcontroller

import (
	"AlekseyMartunov/internal/orders"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

var ErrTooManyRequests = errors.New("to many requests to server")

const tooManyRequests = "429"
const delay = 5

type OrderResponse struct {
	Number  string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual,omitempty"`
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
					return
				}

				orderChan <- r.tryToSendRequest(number)
			}
		}
	}()
	return orderChan
}

func (r *RequestAccrual) get(number string) (orders.Order, error) {
	o := OrderResponse{}
	order := orders.Order{}

	client := resty.New()

	url := strings.Join([]string{r.host, r.url, number}, "")

	resp, err := client.R().
		Get(url)

	if err != nil {
		r.log.Error(err.Error())
		return order, err
	}

	r.log.Warn(resp.Status())
	r.log.Warn(string(resp.Body()))

	if resp.Status() == tooManyRequests {
		return order, ErrTooManyRequests
	}

	err = json.Unmarshal(resp.Body(), &o)
	if err != nil {
		r.log.Error(err.Error())
		return order, err
	}

	order.Number = o.Number
	order.Status = o.Status
	order.Accrual = o.Accrual

	return order, nil
}

func (r *RequestAccrual) tryToSendRequest(number string) orders.Order {
	for {
		ord, err := r.get(number)
		if errors.Is(err, ErrTooManyRequests) {
			time.Sleep(delay * time.Second)
			continue
		}
		return ord
	}
}
