package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5"
)

func (us *UserStorage) GetIDByUUID(ctx context.Context, uuid string) (int, error) {
	query := `SELECT client_id FROM client WHERE client_uuid = $1`
	res := us.conn.QueryRow(ctx, query, uuid)

	var id sql.NullInt64

	err := res.Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return -1, ErrUserDoseNotExist
		}
		return -1, err
	}

	return int(id.Int64), nil
}
