package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (us *UserStorage) Register(ctx context.Context, login, password string) error {
	query := `INSERT INTO client (login, password) VALUES ($1, $2)`

	_, err := us.conn.Exec(ctx, query, login, password)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return LoginAlreadyUsedErr
		}
		return err
	}
	return nil
}
