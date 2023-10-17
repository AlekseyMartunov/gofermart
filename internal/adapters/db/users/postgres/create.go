package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (us *UserStorage) Create(ctx context.Context, login, password string) (string, error) {
	query := `INSERT INTO client (login, password) VALUES ($1, $2) RETURNING client_uuid`

	res := us.conn.QueryRow(ctx, query, login, password)
	var uuid string

	err := res.Scan(&uuid)

	if err != nil {
		fmt.Println("-----------------------")
		fmt.Println("err")
		fmt.Println("-----------------------")
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return "", ErrLoginAlreadyUsed
		}
		return "", err
	}
	return uuid, nil
}
