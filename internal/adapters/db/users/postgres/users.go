package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type UserStorage struct {
	conn *pgx.Conn
	log  Logger
}

func NewUserStorage(conn *pgx.Conn, l Logger) *UserStorage {
	return &UserStorage{
		conn: conn,
		log:  l,
	}
}

func (us *UserStorage) Register(ctx context.Context, login, password string) error {
	query := `INSERT INTO client (login, password) VALUES ($1, $2)`

	res, _ := us.conn.Exec(ctx, query, login, password)
	fmt.Println(res.String())
	return nil
}
