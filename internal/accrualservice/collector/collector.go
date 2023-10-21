package collector

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type Collector struct {
	conn *pgxpool.Pool
	log  logger
}

func NewCollector(conn *pgxpool.Pool, l logger) *Collector {
	return &Collector{
		conn: conn,
		log:  l,
	}
}

func (c *Collector) Run(ctx context.Context, t time.Duration) chan string {
	ticker := time.NewTicker(t)
	ch := make(chan string)

	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				c.log.Warn("Stop collect the data")
				return

			case <-ticker.C:
				c.collect(ctx, ch)
			}
		}
	}()

	return ch
}

func (c *Collector) collect(ctx context.Context, ch chan string) {
	query := `SELECT order_number  
				FROM orders INNER JOIN status ON orders.fk_order_status = status.status_id
				WHERE status.status_name IN ('NEW', 'PROCESSING')`

	rows, err := c.conn.Query(ctx, query)
	if err != nil {
		c.log.Error(err.Error())
	}

	for rows.Next() {
		var number string
		err = rows.Scan(&number)
		if err != nil {
			c.log.Error(err.Error())
			continue
		}
		ch <- number
	}
}
