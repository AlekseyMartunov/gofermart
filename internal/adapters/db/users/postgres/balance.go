package postgres

import (
	"context"
	"github.com/shopspring/decimal"

	"AlekseyMartunov/internal/users"
)

func (us *UserStorage) Balance(ctx context.Context, UserID int) (users.User, error) {
	query := `SELECT bonuses, withdrawn FROM user_balance
 				WHERE fk_user_id = $1`

	row := us.conn.QueryRow(ctx, query, UserID)
	u := users.User{}

	x := decimal.Decimal{}
	y := decimal.Decimal{}

	err := row.Scan(&x, &y)

	if err != nil {
		return u, err
	}

	u.Bonuses = x.InexactFloat64()
	u.Withdrawn = y.InexactFloat64()

	return u, nil
}
