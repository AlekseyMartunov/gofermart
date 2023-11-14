package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"AlekseyMartunov/internal/orders"
)

func (os *OrderStorage) Create(ctx context.Context, order orders.Order) error {
	query := `INSERT INTO orders (order_number, created_time, fk_user_id)
 				VALUES ($1, $2, $3)`

	_, err := os.conn.Exec(ctx, query, order.Number, order.CreatedTime, order.UserID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return ErrOrderAlreadyCreated
		}
		os.log.Error(err.Error())
		return err
	}

	return nil
}
