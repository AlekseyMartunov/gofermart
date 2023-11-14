package postgres

import (
	"context"
	"github.com/shopspring/decimal"

	"AlekseyMartunov/internal/users"
)

func (us *UserStorage) Balance(ctx context.Context, UserID int) (users.User, error) {
	query := `SELECT bonuses, withdrawn FROM client
  				WHERE client_id = $1`

	row := us.conn.QueryRow(ctx, query, UserID)

	x := decimal.Decimal{}
	y := decimal.Decimal{}

	err := row.Scan(&x, &y)

	u := users.User{}
	if err != nil {
		return u, err
	}

	u.Bonuses = x.InexactFloat64()
	u.Withdrawn = y.InexactFloat64()

	return u, nil
}
