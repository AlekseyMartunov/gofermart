package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"AlekseyMartunov/internal/orders"
)

func (os *OrderStorage) AddDiscount(ctx context.Context, order orders.Order) error {
	query1 := `INSERT INTO orders (order_number, created_time, fk_user_id)
 				VALUES ($1, $2, $3) RETURNING order_id`

	query2 := `UPDATE orders SET 
                   discount = $1
                   WHERE order_id = $2;`

	query3 := `UPDATE client SET 
				   bonuses = bonuses - $1,
				   withdrawn = withdrawn + $1
				   WHERE client_id = $2;`

	query4 := `INSERT INTO history (order_id, amount, fk_user_id) VALUES 
                        ($1, $2, $3)`

	txOpt := pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadWrite,
		DeferrableMode: pgx.Deferrable,
	}

	tx, err := os.conn.BeginTx(ctx, txOpt)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, query1, order.Number, order.CreatedTime, order.UserID)
	var orderID int
	err = row.Scan(&orderID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, query2, order.Discount, orderID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, query3, order.Discount, order.UserID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			tx.Rollback(ctx)
			return ErrNotEnoughMoney
		}
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, query4, orderID, order.Discount, order.UserID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
