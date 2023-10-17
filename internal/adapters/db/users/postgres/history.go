package postgres

import (
	"AlekseyMartunov/internal/users"
	"context"
	"time"
)

func (us *UserStorage) GetHistory(ctx context.Context, userID int) ([]users.HistoryElement, error) {
	query := `SELECT orders.order_number, history.created_time, amount
				FROM history INNER JOIN orders ON orders.order_id = history.order_id
				WHERE history.fk_user_id = $1
				ORDER BY created_time DESC;`

	rows, err := us.conn.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	res := make([]users.HistoryElement, 0, 20)

	for rows.Next() {
		var number string
		var amount float64
		var createdTime time.Time

		err := rows.Scan(&number, &createdTime, &amount)
		if err != nil {
			return nil, err
		}

		h := users.HistoryElement{
			Order:        number,
			Amount:       amount,
			WriteOffTime: createdTime,
		}

		res = append(res, h)
	}

	if len(res) == 0 {
		return nil, ErrEmptyHistory
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}
