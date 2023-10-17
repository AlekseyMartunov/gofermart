package postgres

import (
	"errors"
	"github.com/jackc/pgx/v5"
)

var ErrOrderAlreadyCreated = errors.New("the order was already loaded")
var ErrEmptyResult = errors.New("empty result")

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type OrderStorage struct {
	conn *pgx.Conn
	log  Logger
}

func NewOrderStorage(conn *pgx.Conn, l Logger) *OrderStorage {
	return &OrderStorage{
		conn: conn,
		log:  l,
	}
}
