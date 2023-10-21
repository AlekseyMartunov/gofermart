package postgres

import (
	"AlekseyMartunov/internal/orders"
	"context"
)

func (os *OrderStorage) Update(ctx context.Context, orders ...orders.Order) error {
	query := `UPDATE orders 
				SET fk_order_status = (SELECT status_id FROM status 
						WHERE status_name = $1),
					accrual = $2
					FROM orders AS O INNER JOIN status AS S ON O.fk_order_status = S.status_id
					WHERE orders.order_number = $3;`

	for _, o := range orders {
		_, err := os.conn.Exec(ctx, query, o.Status, o.Accrual, o.Number)
		if err != nil {
			return err
		}
	}

	return nil

}
