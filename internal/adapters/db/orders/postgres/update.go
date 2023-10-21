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
				FROM orders AS O 
					INNER JOIN status AS S ON O.fk_order_status = S.status_id
				WHERE orders.order_number = $3;`

	query2 := `UPDATE client 
				SET bonuses = bonuses +  $1
				WHERE client_id = (SELECT U.client_id FROM orders AS O
									INNER JOIN client AS U ON O.fk_user_id = U.client_id
									WHERE O.order_number = $2);`

	for _, o := range orders {
		_, err := os.conn.Exec(ctx, query, o.Status, o.Accrual, o.Number)
		if err != nil {
			return err
		}
		_, err = os.conn.Exec(ctx, query2, o.Accrual, o.Number)
		if err != nil {
			return err
		}
	}

	return nil

}
