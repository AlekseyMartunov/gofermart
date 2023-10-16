package postgres

import (
	"context"
	"database/sql"
)

func (os *OrderStorage) GetUserID(ctx context.Context, number string) (int, error) {
	query := `SELECT fk_user_id FROM "order" WHERE order_number = $1`

	res := os.conn.QueryRow(ctx, query, number)

	var userID sql.NullInt64

	err := res.Scan(&userID)

	if userID.Valid {
		return int(userID.Int64), nil
	}

	return -1, err
}
