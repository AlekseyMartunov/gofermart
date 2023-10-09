package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (us *UserStorage) CheckUser(ctx context.Context, login, password string) (string, error) {
	query := `SELECT client_uuid FROM client
				WHERE login = $1 AND password = $2`

	res := us.conn.QueryRow(ctx, query, login, password)

	var uuid sql.NullString
	err := res.Scan(&uuid)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.String, ErrWrongLoginOrPassword
		}
		return uuid.String, err
	}

	return uuid.String, err
}
