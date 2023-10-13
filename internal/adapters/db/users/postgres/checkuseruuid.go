package postgres

import (
	"context"
)

func (us *UserStorage) CheckUserUUID(ctx context.Context, userUUID string) error {
	query := `SELECT EXISTS(SELECT 1 FROM client WHERE client_uuid::text =$1)`
	res := us.conn.QueryRow(ctx, query, userUUID)

	var isExists bool

	err := res.Scan(&isExists)

	if err != nil {
		return err
	}

	if isExists {
		return nil
	}

	return ErrUserDoseNotExist

}
