package postgres

import (
	"AlekseyMartunov/internal/orders"
	"context"
	"database/sql"
	"time"
)

func (os *OrderStorage) GetOrders(ctx context.Context, userID int) ([]orders.Order, error) {
	query := `SELECT order_number, status.status_name, accrual, created_time
				FROM "order" AS t1 INNER JOIN status ON t1.fk_order_status = status.status_id
				WHERE t1.fk_user_id = $1
				ORDER BY created_time DESC;`

	rows, err := os.conn.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	resOrders := make([]orders.Order, 0, 20)

	for rows.Next() {
		var number string
		var status string
		var accrual sql.NullInt64
		var createdTime time.Time

		err := rows.Scan(&number, &status, &accrual, &createdTime)
		if err != nil {
			return nil, err
		}

		o := orders.Order{
			Number:      number,
			Status:      status,
			Accrual:     int(accrual.Int64),
			CreatedTime: createdTime,
		}
		resOrders = append(resOrders, o)
	}
	if len(resOrders) == 0 {
		return nil, ErrEmptyResult
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return resOrders, nil

}
